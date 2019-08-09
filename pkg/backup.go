package pkg

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/appscode/go/flags"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	appcatalog_cs "kmodules.xyz/custom-resources/client/clientset/versioned"
	kubedbconfig_api "kubedb.dev/apimachinery/apis/config/v1alpha1"
	"stash.appscode.dev/stash/pkg/restic"
	"stash.appscode.dev/stash/pkg/util"
)

const (
	jobMySqlBackup = "stash-xtradb-backup"

	mySqlUser     = "username"
	mySqlPassword = "password"

	mySqlDumpFile        = "dumpfile.sql"
	xtraBackupStreamFile = "xtrabackup.stream"

	mySqlDumpCMD     = "mysqldump"
	mySqlRestoreCMD  = "mysql"
	envMySqlPassword = "MYSQL_PWD"
	defaultDumpArgs  = "--all-databases"
)

func NewCmdBackup() *cobra.Command {
	var (
		masterURL      string
		kubeconfigPath string
		namespace      string
		appBindingName string
		outputDir      string
		dumpArgs       string
		garbdCnf       kubedbconfig_api.GaleraArbitratorConfiguration
		socatRetry     int32

		setupOpt = restic.SetupOptions{
			ScratchDir:  restic.DefaultScratchDir,
			EnableCache: false,
		}
		backupOpt = restic.BackupOptions{
			Host: restic.DefaultHost,
		}
		metrics = restic.MetricsOptions{
			JobName: jobMySqlBackup,
		}
	)

	cmd := &cobra.Command{
		Use:               "backup-percona-xtradb",
		Short:             "Takes a backup of XtraDB DB",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			flags.EnsureRequiredFlags(cmd, "app-binding", "provider", "secret-dir")

			// apply nice, ionice settings from env
			var err error
			setupOpt.Nice, err = util.NiceSettingsFromEnv()
			if err != nil {
				return util.HandleResticError(outputDir, restic.DefaultOutputFileName, err)
			}
			setupOpt.IONice, err = util.IONiceSettingsFromEnv()
			if err != nil {
				return util.HandleResticError(outputDir, restic.DefaultOutputFileName, err)
			}

			// prepare client
			config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
			if err != nil {
				return err
			}
			kubeClient, err := kubernetes.NewForConfig(config)
			if err != nil {
				return err
			}
			appCatalogClient, err := appcatalog_cs.NewForConfig(config)
			if err != nil {
				return err
			}

			// get app binding
			appBinding, err := appCatalogClient.AppcatalogV1alpha1().AppBindings(namespace).Get(appBindingName, metav1.GetOptions{})
			if err != nil {
				return err
			}
			// get secret
			appBindingSecret, err := kubeClient.CoreV1().Secrets(namespace).Get(appBinding.Spec.Secret.Name, metav1.GetOptions{})
			if err != nil {
				return err
			}

			// galera arbitrator config
			if appBinding.Spec.Parameters != nil {
				err = json.Unmarshal(appBinding.Spec.Parameters.Raw, &garbdCnf)
				if err != nil {
					return err
				}
			}

			// init restic wrapper
			resticWrapper, err := restic.NewResticWrapper(setupOpt)
			if err != nil {
				return err
			}

			if garbdCnf.Group == "" { // standalone database
				backupOpt.StdinFileName = mySqlDumpFile

				// set env for mysqlbdump
				resticWrapper.SetEnv(envMySqlPassword, string(appBindingSecret.Data[mySqlPassword]))
				// setup pipe command
				backupOpt.StdinPipeCommand = restic.Command{
					Name: mySqlDumpCMD,
					Args: []interface{}{
						"-u", string(appBindingSecret.Data[mySqlUser]),
						"-h", appBinding.Spec.ClientConfig.Service.Name,
					},
				}
				if dumpArgs != "" {
					backupOpt.StdinPipeCommand.Args = append(backupOpt.StdinPipeCommand.Args, dumpArgs)
				}
			} else { // clustered database
				backupOpt.StdinFileName = xtraBackupStreamFile

				sstRequestOpts := fmt.Sprintf("%s,%d,%s",
					garbdCnf.SSTMethod,
					kubedbconfig_api.GarbdListenPort,
					kubedbconfig_api.GarbdXtrabackupSSTRequestSuffix,
				)

				backupOpt.StdinPipeCommand = restic.Command{
					Name: "bash",
					Args: []interface{}{
						"-c",
						fmt.Sprintf("/backup.sh %s %s %s %s %s",
							garbdCnf.ClusterAddressWithListenOption(),
							garbdCnf.Group,
							sstRequestOpts,
							kubedbconfig_api.GarbdLogFile,
							kubedbconfig_api.SOCATOption(socatRetry),
						),
					},
				}
			}

			// wait for DB ready
			waitForDBReady(appBinding.Spec.ClientConfig.Service.Name, appBinding.Spec.ClientConfig.Service.Port)

			// Run backup
			backupOutput, backupErr := resticWrapper.RunBackup(backupOpt)
			// If metrics are enabled then generate metrics
			if metrics.Enabled {
				err := backupOutput.HandleMetrics(&metrics, backupErr)
				if err != nil {
					return errors.NewAggregate([]error{backupErr, err})
				}
			}
			// If output directory specified, then write the output in "output.json" file in the specified directory
			if backupErr == nil && outputDir != "" {
				err := backupOutput.WriteOutput(filepath.Join(outputDir, restic.DefaultOutputFileName))
				if err != nil {
					return err
				}
			}
			return backupErr
		},
	}

	cmd.Flags().StringVar(&dumpArgs, "xtradb-args", defaultDumpArgs, "Additional arguments")
	cmd.Flags().Int32Var(&socatRetry, "socat-retry", kubedbconfig_api.SOCATOptionRetry, "Additional arguments")

	cmd.Flags().StringVar(&masterURL, "master", masterURL, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	cmd.Flags().StringVar(&kubeconfigPath, "kubeconfig", kubeconfigPath, "Path to kubeconfig file with authorization information (the master location is set by the master flag).")
	cmd.Flags().StringVar(&namespace, "namespace", "default", "Namespace of Backup/Restore Session")
	cmd.Flags().StringVar(&appBindingName, "app-binding", appBindingName, "Name of the app binding")

	cmd.Flags().StringVar(&setupOpt.Provider, "provider", setupOpt.Provider, "Backend provider (i.e. gcs, s3, azure etc)")
	cmd.Flags().StringVar(&setupOpt.Bucket, "bucket", setupOpt.Bucket, "Name of the cloud bucket/container (keep empty for local backend)")
	cmd.Flags().StringVar(&setupOpt.Endpoint, "endpoint", setupOpt.Endpoint, "Endpoint for s3/s3 compatible backend")
	cmd.Flags().StringVar(&setupOpt.URL, "rest-server-url", setupOpt.URL, "URL for rest backend")
	cmd.Flags().StringVar(&setupOpt.Path, "path", setupOpt.Path, "Directory inside the bucket where backup will be stored")
	cmd.Flags().StringVar(&setupOpt.SecretDir, "secret-dir", setupOpt.SecretDir, "Directory where storage secret has been mounted")
	cmd.Flags().StringVar(&setupOpt.ScratchDir, "scratch-dir", setupOpt.ScratchDir, "Temporary directory")
	cmd.Flags().BoolVar(&setupOpt.EnableCache, "enable-cache", setupOpt.EnableCache, "Specify whether to enable caching for restic")
	cmd.Flags().IntVar(&setupOpt.MaxConnections, "max-connections", setupOpt.MaxConnections, "Specify maximum concurrent connections for GCS, Azure and B2 backend")

	cmd.Flags().StringVar(&backupOpt.Host, "hostname", backupOpt.Host, "Name of the host machine")

	cmd.Flags().IntVar(&backupOpt.RetentionPolicy.KeepLast, "retention-keep-last", backupOpt.RetentionPolicy.KeepLast, "Specify value for retention strategy")
	cmd.Flags().IntVar(&backupOpt.RetentionPolicy.KeepHourly, "retention-keep-hourly", backupOpt.RetentionPolicy.KeepHourly, "Specify value for retention strategy")
	cmd.Flags().IntVar(&backupOpt.RetentionPolicy.KeepDaily, "retention-keep-daily", backupOpt.RetentionPolicy.KeepDaily, "Specify value for retention strategy")
	cmd.Flags().IntVar(&backupOpt.RetentionPolicy.KeepWeekly, "retention-keep-weekly", backupOpt.RetentionPolicy.KeepWeekly, "Specify value for retention strategy")
	cmd.Flags().IntVar(&backupOpt.RetentionPolicy.KeepMonthly, "retention-keep-monthly", backupOpt.RetentionPolicy.KeepMonthly, "Specify value for retention strategy")
	cmd.Flags().IntVar(&backupOpt.RetentionPolicy.KeepYearly, "retention-keep-yearly", backupOpt.RetentionPolicy.KeepYearly, "Specify value for retention strategy")
	cmd.Flags().StringSliceVar(&backupOpt.RetentionPolicy.KeepTags, "retention-keep-tags", backupOpt.RetentionPolicy.KeepTags, "Specify value for retention strategy")
	cmd.Flags().BoolVar(&backupOpt.RetentionPolicy.Prune, "retention-prune", backupOpt.RetentionPolicy.Prune, "Specify whether to prune old snapshot data")
	cmd.Flags().BoolVar(&backupOpt.RetentionPolicy.DryRun, "retention-dry-run", backupOpt.RetentionPolicy.DryRun, "Specify whether to test retention policy without deleting actual data")

	cmd.Flags().StringVar(&outputDir, "output-dir", outputDir, "Directory where output.json file will be written (keep empty if you don't need to write output in file)")

	cmd.Flags().BoolVar(&metrics.Enabled, "metrics-enabled", metrics.Enabled, "Specify whether to export Prometheus metrics")
	cmd.Flags().StringVar(&metrics.PushgatewayURL, "metrics-pushgateway-url", metrics.PushgatewayURL, "Pushgateway URL where the metrics will be pushed")
	cmd.Flags().StringVar(&metrics.MetricFileDir, "metrics-dir", metrics.MetricFileDir, "Directory where to write metric.prom file (keep empty if you don't want to write metric in a text file)")
	cmd.Flags().StringSliceVar(&metrics.Labels, "metrics-labels", metrics.Labels, "Labels to apply in exported metrics")

	return cmd
}

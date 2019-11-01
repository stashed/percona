/*
Copyright The Stash Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package pkg

import (
	"path/filepath"
	"strings"

	api_v1beta1 "stash.appscode.dev/stash/apis/stash/v1beta1"
	"stash.appscode.dev/stash/pkg/restic"
	"stash.appscode.dev/stash/pkg/util"

	"github.com/appscode/go/flags"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	appcatalog_cs "kmodules.xyz/custom-resources/client/clientset/versioned"
)

func NewCmdRestore() *cobra.Command {
	var (
		masterURL      string
		kubeconfigPath string
		opt            = perconaOptions{
			setupOptions: restic.SetupOptions{
				ScratchDir:  restic.DefaultScratchDir,
				EnableCache: false,
			},
			dumpOptions: restic.DumpOptions{
				Host: restic.DefaultHost,
			},
		}
	)

	cmd := &cobra.Command{
		Use:               "restore-percona-xtradb",
		Short:             "Restores XtraDB DB Backup",
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			flags.EnsureRequiredFlags(cmd, "appbinding", "provider", "secret-dir")

			// prepare client
			config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
			if err != nil {
				return err
			}
			opt.kubeClient, err = kubernetes.NewForConfig(config)
			if err != nil {
				return err
			}
			opt.catalogClient, err = appcatalog_cs.NewForConfig(config)
			if err != nil {
				return err
			}

			var restoreOutput *restic.RestoreOutput
			restoreOutput, err = opt.restorePerconaXtraDB()
			if err != nil {
				restoreOutput = &restic.RestoreOutput{
					HostRestoreStats: []api_v1beta1.HostRestoreStats{
						{
							Hostname: opt.dumpOptions.Host,
							Phase:    api_v1beta1.HostRestoreFailed,
							Error:    err.Error(),
						},
					},
				}
			}
			// If output directory specified, then write the output in "output.json" file in the specified directory
			if opt.outputDir != "" {
				return restoreOutput.WriteOutput(filepath.Join(opt.outputDir, restic.DefaultOutputFileName))
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&opt.xtradbArgs, "xtradb-args", opt.xtradbArgs, "Additional arguments")
	cmd.Flags().Int32Var(&opt.targetAppReplicas, "target-app-replicas", opt.targetAppReplicas, "Additional arguments")

	cmd.Flags().StringVar(&masterURL, "master", masterURL, "The address of the Kubernetes API server (overrides any value in kubeconfig)")
	cmd.Flags().StringVar(&kubeconfigPath, "kubeconfig", kubeconfigPath, "Path to kubeconfig file with authorization information (the master location is set by the master flag).")
	cmd.Flags().StringVar(&opt.namespace, "namespace", "default", "Namespace of Backup/Restore Session")
	cmd.Flags().StringVar(&opt.appBindingName, "appbinding", opt.appBindingName, "Name of the app binding")

	cmd.Flags().StringVar(&opt.setupOptions.Provider, "provider", opt.setupOptions.Provider, "Backend provider (i.e. gcs, s3, azure etc)")
	cmd.Flags().StringVar(&opt.setupOptions.Bucket, "bucket", opt.setupOptions.Bucket, "Name of the cloud bucket/container (keep empty for local backend)")
	cmd.Flags().StringVar(&opt.setupOptions.Endpoint, "endpoint", opt.setupOptions.Endpoint, "Endpoint for s3/s3 compatible backend or REST server URL")
	cmd.Flags().StringVar(&opt.setupOptions.Path, "path", opt.setupOptions.Path, "Directory inside the bucket where backup will be stored")
	cmd.Flags().StringVar(&opt.setupOptions.SecretDir, "secret-dir", opt.setupOptions.SecretDir, "Directory where storage secret has been mounted")
	cmd.Flags().StringVar(&opt.setupOptions.ScratchDir, "scratch-dir", opt.setupOptions.ScratchDir, "Temporary directory")
	cmd.Flags().BoolVar(&opt.setupOptions.EnableCache, "enable-cache", opt.setupOptions.EnableCache, "Specify whether to enable caching for restic")
	cmd.Flags().IntVar(&opt.setupOptions.MaxConnections, "max-connections", opt.setupOptions.MaxConnections, "Specify maximum concurrent connections for GCS, Azure and B2 backend")

	cmd.Flags().StringVar(&opt.dumpOptions.Host, "hostname", opt.dumpOptions.Host, "Name of the host machine")
	cmd.Flags().StringVar(&opt.dumpOptions.SourceHost, "source-hostname", opt.dumpOptions.SourceHost, "Name of the host from where data will be restored")
	// TODO: sliceVar
	cmd.Flags().StringVar(&opt.dumpOptions.Snapshot, "snapshot", opt.dumpOptions.Snapshot, "Snapshot to dump")

	cmd.Flags().StringVar(&opt.outputDir, "output-dir", opt.outputDir, "Directory where output.json file will be written (keep empty if you don't need to write output in file)")

	return cmd
}

func (opt *perconaOptions) restorePerconaXtraDB() (*restic.RestoreOutput, error) {
	// apply nice, ionice settings from env
	var err error
	opt.setupOptions.Nice, err = util.NiceSettingsFromEnv()
	if err != nil {
		return nil, err
	}
	opt.setupOptions.IONice, err = util.IONiceSettingsFromEnv()
	if err != nil {
		return nil, err
	}

	// init restic wrapper
	resticWrapper, err := restic.NewResticWrapper(opt.setupOptions)
	if err != nil {
		return nil, err
	}

	if opt.targetAppReplicas == 1 {
		// get app binding
		appBinding, err := opt.catalogClient.AppcatalogV1alpha1().AppBindings(opt.namespace).Get(opt.appBindingName, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		// get secret
		appBindingSecret, err := opt.kubeClient.CoreV1().Secrets(opt.namespace).Get(appBinding.Spec.Secret.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}

		// set env for mysql
		resticWrapper.SetEnv(envMySqlPassword, string(appBindingSecret.Data[mySqlPassword]))

		// set backed up file name
		opt.dumpOptions.FileName = mySqlDumpFile

		// setup pipe command
		opt.dumpOptions.StdoutPipeCommand = restic.Command{
			Name: mySqlRestoreCMD,
			Args: []interface{}{
				"-u", string(appBindingSecret.Data[mySqlUser]),
				"-h", appBinding.Spec.ClientConfig.Service.Name,
			},
		}
		for _, arg := range strings.Fields(opt.xtradbArgs) {
			opt.dumpOptions.StdoutPipeCommand.Args = append(opt.dumpOptions.StdoutPipeCommand.Args, arg)
		}
		// wait for DB ready
		waitForDBReady(appBinding.Spec.ClientConfig.Service.Name, appBinding.Spec.ClientConfig.Service.Port)
	} else {
		// set backed up file name
		opt.dumpOptions.FileName = xtraBackupStreamFile

		// setup pipe command
		opt.dumpOptions.StdoutPipeCommand = restic.Command{
			Name: "bash",
			Args: []interface{}{
				"-c",
				"/restore.sh /var/lib/mysql",
			},
		}
	}

	// Run dump
	return resticWrapper.Dump(opt.dumpOptions)
}

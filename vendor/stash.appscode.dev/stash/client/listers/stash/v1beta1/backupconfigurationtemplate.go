/*
Copyright 2019 The Stash Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1beta1 "stash.appscode.dev/stash/apis/stash/v1beta1"
)

// BackupConfigurationTemplateLister helps list BackupConfigurationTemplates.
type BackupConfigurationTemplateLister interface {
	// List lists all BackupConfigurationTemplates in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.BackupConfigurationTemplate, err error)
	// Get retrieves the BackupConfigurationTemplate from the index for a given name.
	Get(name string) (*v1beta1.BackupConfigurationTemplate, error)
	BackupConfigurationTemplateListerExpansion
}

// backupConfigurationTemplateLister implements the BackupConfigurationTemplateLister interface.
type backupConfigurationTemplateLister struct {
	indexer cache.Indexer
}

// NewBackupConfigurationTemplateLister returns a new BackupConfigurationTemplateLister.
func NewBackupConfigurationTemplateLister(indexer cache.Indexer) BackupConfigurationTemplateLister {
	return &backupConfigurationTemplateLister{indexer: indexer}
}

// List lists all BackupConfigurationTemplates in the indexer.
func (s *backupConfigurationTemplateLister) List(selector labels.Selector) (ret []*v1beta1.BackupConfigurationTemplate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.BackupConfigurationTemplate))
	})
	return ret, err
}

// Get retrieves the BackupConfigurationTemplate from the index for a given name.
func (s *backupConfigurationTemplateLister) Get(name string) (*v1beta1.BackupConfigurationTemplate, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("backupconfigurationtemplate"), name)
	}
	return obj.(*v1beta1.BackupConfigurationTemplate), nil
}

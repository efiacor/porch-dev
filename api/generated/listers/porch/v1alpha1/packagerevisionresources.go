// Copyright 2023 The kpt and Nephio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/GoogleContainerTools/kpt/porch/api/porch/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PackageRevisionResourcesLister helps list PackageRevisionResourceses.
// All objects returned here must be treated as read-only.
type PackageRevisionResourcesLister interface {
	// List lists all PackageRevisionResourceses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PackageRevisionResources, err error)
	// PackageRevisionResourceses returns an object that can list and get PackageRevisionResourceses.
	PackageRevisionResourceses(namespace string) PackageRevisionResourcesNamespaceLister
	PackageRevisionResourcesListerExpansion
}

// packageRevisionResourcesLister implements the PackageRevisionResourcesLister interface.
type packageRevisionResourcesLister struct {
	indexer cache.Indexer
}

// NewPackageRevisionResourcesLister returns a new PackageRevisionResourcesLister.
func NewPackageRevisionResourcesLister(indexer cache.Indexer) PackageRevisionResourcesLister {
	return &packageRevisionResourcesLister{indexer: indexer}
}

// List lists all PackageRevisionResourceses in the indexer.
func (s *packageRevisionResourcesLister) List(selector labels.Selector) (ret []*v1alpha1.PackageRevisionResources, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PackageRevisionResources))
	})
	return ret, err
}

// PackageRevisionResourceses returns an object that can list and get PackageRevisionResourceses.
func (s *packageRevisionResourcesLister) PackageRevisionResourceses(namespace string) PackageRevisionResourcesNamespaceLister {
	return packageRevisionResourcesNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PackageRevisionResourcesNamespaceLister helps list and get PackageRevisionResourceses.
// All objects returned here must be treated as read-only.
type PackageRevisionResourcesNamespaceLister interface {
	// List lists all PackageRevisionResourceses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PackageRevisionResources, err error)
	// Get retrieves the PackageRevisionResources from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PackageRevisionResources, error)
	PackageRevisionResourcesNamespaceListerExpansion
}

// packageRevisionResourcesNamespaceLister implements the PackageRevisionResourcesNamespaceLister
// interface.
type packageRevisionResourcesNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PackageRevisionResourceses in the indexer for a given namespace.
func (s packageRevisionResourcesNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PackageRevisionResources, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PackageRevisionResources))
	})
	return ret, err
}

// Get retrieves the PackageRevisionResources from the indexer for a given namespace and name.
func (s packageRevisionResourcesNamespaceLister) Get(name string) (*v1alpha1.PackageRevisionResources, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("packagerevisionresources"), name)
	}
	return obj.(*v1alpha1.PackageRevisionResources), nil
}
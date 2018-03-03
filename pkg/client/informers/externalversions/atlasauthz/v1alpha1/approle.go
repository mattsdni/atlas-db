/*
Copyright 2018 Infoblox, Inc.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	atlasauthz_v1alpha1 "github.com/infobloxopen/atlas/pkg/apis/atlasauthz/v1alpha1"
	versioned "github.com/infobloxopen/atlas/pkg/client/clientset/versioned"
	internalinterfaces "github.com/infobloxopen/atlas/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/infobloxopen/atlas/pkg/client/listers/atlasauthz/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AppRoleInformer provides access to a shared informer and lister for
// AppRoles.
type AppRoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.AppRoleLister
}

type appRoleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAppRoleInformer constructs a new informer for AppRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAppRoleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAppRoleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAppRoleInformer constructs a new informer for AppRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAppRoleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AtlasauthzV1alpha1().AppRoles(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AtlasauthzV1alpha1().AppRoles(namespace).Watch(options)
			},
		},
		&atlasauthz_v1alpha1.AppRole{},
		resyncPeriod,
		indexers,
	)
}

func (f *appRoleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAppRoleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *appRoleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&atlasauthz_v1alpha1.AppRole{}, f.defaultInformer)
}

func (f *appRoleInformer) Lister() v1alpha1.AppRoleLister {
	return v1alpha1.NewAppRoleLister(f.Informer().GetIndexer())
}

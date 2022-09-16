// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	kubeaninfomanifestv1alpha1 "kubean.io/api/apis/kubeaninfomanifest/v1alpha1"
	versioned "kubean.io/api/generated/kubeaninfomanifest/clientset/versioned"
	internalinterfaces "kubean.io/api/generated/kubeaninfomanifest/informers/externalversions/internalinterfaces"
	v1alpha1 "kubean.io/api/generated/kubeaninfomanifest/listers/kubeaninfomanifest/v1alpha1"
)

// KubeanInfoManifestInformer provides access to a shared informer and lister for
// KubeanInfoManifests.
type KubeanInfoManifestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.KubeanInfoManifestLister
}

type kubeanInfoManifestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKubeanInfoManifestInformer constructs a new informer for KubeanInfoManifest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeanInfoManifestInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeanInfoManifestInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKubeanInfoManifestInformer constructs a new informer for KubeanInfoManifest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeanInfoManifestInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeanV1alpha1().KubeanInfoManifests().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeanV1alpha1().KubeanInfoManifests().Watch(context.TODO(), options)
			},
		},
		&kubeaninfomanifestv1alpha1.KubeanInfoManifest{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeanInfoManifestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeanInfoManifestInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeanInfoManifestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeaninfomanifestv1alpha1.KubeanInfoManifest{}, f.defaultInformer)
}

func (f *kubeanInfoManifestInformer) Lister() v1alpha1.KubeanInfoManifestLister {
	return v1alpha1.NewKubeanInfoManifestLister(f.Informer().GetIndexer())
}
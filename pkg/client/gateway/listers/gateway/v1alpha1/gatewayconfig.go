// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/nginxinc/nginx-gateway-kubernetes/pkg/apis/gateway/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GatewayConfigLister helps list GatewayConfigs.
// All objects returned here must be treated as read-only.
type GatewayConfigLister interface {
	// List lists all GatewayConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GatewayConfig, err error)
	// Get retrieves the GatewayConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GatewayConfig, error)
	GatewayConfigListerExpansion
}

// gatewayConfigLister implements the GatewayConfigLister interface.
type gatewayConfigLister struct {
	indexer cache.Indexer
}

// NewGatewayConfigLister returns a new GatewayConfigLister.
func NewGatewayConfigLister(indexer cache.Indexer) GatewayConfigLister {
	return &gatewayConfigLister{indexer: indexer}
}

// List lists all GatewayConfigs in the indexer.
func (s *gatewayConfigLister) List(selector labels.Selector) (ret []*v1alpha1.GatewayConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GatewayConfig))
	})
	return ret, err
}

// Get retrieves the GatewayConfig from the index for a given name.
func (s *gatewayConfigLister) Get(name string) (*v1alpha1.GatewayConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("gatewayconfig"), name)
	}
	return obj.(*v1alpha1.GatewayConfig), nil
}
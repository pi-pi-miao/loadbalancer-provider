/*
Copyright 2020 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/caicloud/clientset/customclient/scheme"
	v1beta1 "github.com/caicloud/clientset/pkg/apis/workload/v1beta1"
	rest "k8s.io/client-go/rest"
)

type WorkloadV1beta1Interface interface {
	RESTClient() rest.Interface
	WorkloadsGetter
}

// WorkloadV1beta1Client is used to interact with features provided by the workload.caicloud.io group.
type WorkloadV1beta1Client struct {
	restClient rest.Interface
}

func (c *WorkloadV1beta1Client) Workloads(namespace string) WorkloadInterface {
	return newWorkloads(c, namespace)
}

// NewForConfig creates a new WorkloadV1beta1Client for the given config.
func NewForConfig(c *rest.Config) (*WorkloadV1beta1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &WorkloadV1beta1Client{client}, nil
}

// NewForConfigOrDie creates a new WorkloadV1beta1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *WorkloadV1beta1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new WorkloadV1beta1Client for the given RESTClient.
func New(c rest.Interface) *WorkloadV1beta1Client {
	return &WorkloadV1beta1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1beta1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *WorkloadV1beta1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
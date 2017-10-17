package platform

import (
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

//Kubernetes ...
type Kubernetes struct {
	clientset *kubernetes.Clientset
	config    *rest.Config
}

//NewKubernetes object
func NewKubernetes() (*Kubernetes, error) {
	//InCluster...
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &Kubernetes{clientset: clientset, config: config}, nil
}

//CreateNamespace within kubernetes
func (k *Kubernetes) CreateNamespace(namespace string) error {
	ns := &v1.Namespace{}
	ns.SetNamespace(namespace)
	_, err := k.clientset.CoreV1().Namespaces().Create(ns)
	return err
}

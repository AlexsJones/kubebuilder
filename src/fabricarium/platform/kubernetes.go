package platform

import (
	"os"
	"path/filepath"

	"github.com/AlexsJones/kubebuilder/src/log"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//Kubernetes ...
type Kubernetes struct {
	clientset *kubernetes.Clientset
	config    *rest.Config
}

//NewKubernetes object
func NewKubernetes(inclusterConfig bool) (*Kubernetes, error) {

	//InCluster...
	var config *rest.Config
	var err error
	if inclusterConfig {
		logger.GetInstance().Log("Using in cluster configuration")
		config, err = rest.InClusterConfig()
		if err != nil {
			return nil, err
		}
	} else {
		logger.GetInstance().Log("Using out of cluster configuration")
		kubeconfig := filepath.Join(func() string {
			if h := os.Getenv("HOME"); h != "" {
				return h
			}
			return os.Getenv("USERPROFILE")
		}(), ".kube", "config")
		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
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
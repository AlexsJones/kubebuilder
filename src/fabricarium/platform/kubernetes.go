package platform

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"

	"github.com/AlexsJones/kubebuilder/src/data"
	"github.com/AlexsJones/kubebuilder/src/log"
	"k8s.io/api/core/v1"
	beta "k8s.io/api/extensions/v1beta1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	//This is required for gcp auth provider scope
)

//Kubernetes ...
type Kubernetes struct {
	clientset *kubernetes.Clientset
	config    *rest.Config
}

//NewKubernetes object
func NewKubernetes(masterURL string, inclusterConfig bool) (*Kubernetes, error) {

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
		config, err = clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
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

//ValidateDeployment from deserialisation of YAML
func (k *Kubernetes) ValidateDeployment(build *data.BuildDefinition) (bool, error) {

	logger.GetInstance().Log("Attempting serialisation of YAML")

	d := scheme.Codecs.UniversalDeserializer()
	_, _, err := d.Decode([]byte(build.Kubernetes.Deployment), nil, nil)
	if err != nil {
		logger.GetInstance().Log(fmt.Sprintf("could not decode yaml: %s\n%s", build.Kubernetes.Deployment, err))
		return false, err
	}
	logger.GetInstance().Log("Deployment YAML okay")
	return true, nil
}

//CreateNamespace within kubernetes
func (k *Kubernetes) CreateNamespace(namespace string) (*v1.Namespace, error) {
	if ns, err := k.GetNamespace(namespace); err == nil {
		logger.GetInstance().Log("Found existing namespace")
		return ns, err
	}
	ns := &v1.Namespace{}

	ns.SetName(namespace)

	ns, err := k.clientset.CoreV1().Namespaces().Create(ns)
	return ns, err
}

//GetNamespace within kubernetes
func (k *Kubernetes) GetNamespace(namespace string) (*v1.Namespace, error) {

	ns, err := k.clientset.CoreV1().Namespaces().Get(namespace, meta.GetOptions{})
	return ns, err
}

//CreateDeployment ...
func (k *Kubernetes) CreateDeployment(build *data.BuildDefinition) error {

	deserializer := serializer.NewCodecFactory(clientsetscheme.Scheme).UniversalDeserializer()
	obj, _, err := deserializer.Decode([]byte(build.Kubernetes.Deployment), nil, nil)

	if err != nil {
		logger.GetInstance().Log(fmt.Sprintf("could not decode yaml: %s\n%s", build.Kubernetes.Deployment, err))
		return err
	}
	logger.GetInstance().Info(reflect.TypeOf(obj).String())

	logger.GetInstance().Info("Successfully serialised deployment!")

	deploymentClient := k.clientset.ExtensionsV1beta1().Deployments(build.Kubernetes.Namespace)

	_, err = deploymentClient.Create(obj.(*beta.Deployment))
	if err != nil {
		logger.GetInstance().Log("Trying to update existing deployment....")

		_, err := deploymentClient.Update(obj.(*beta.Deployment))
		if err == nil {
			logger.GetInstance().Info("Updated existing deployment")
			return nil
		}

		return err
	}
	logger.GetInstance().Info("Deployed!")
	return nil
}

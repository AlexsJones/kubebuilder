package platform

import (
	"github.com/AlexsJones/kubebuilder/src/data"
	"k8s.io/api/core/v1"
	beta "k8s.io/api/extensions/v1beta1"
)

//IPlatform interface for container platform
type IPlatform interface {
	ValidateDeployment(build *data.BuildDefinition) (bool, error)
	CreateNamespace(string) (*v1.Namespace, error)
	CreateDeployment(build *data.BuildDefinition) (*beta.Deployment, error)
	CreateService(build *data.BuildDefinition) error
}

//ValidateDeployment from deserialisation of YAML
func ValidateDeployment(i IPlatform, build *data.BuildDefinition) (bool, error) {
	return i.ValidateDeployment(build)
}

//CreateNamespace within the platform cluster
func CreateNamespace(i IPlatform, namespace string) (*v1.Namespace, error) {
	return i.CreateNamespace(namespace)
}

//CreateDeployment within the platform cluster
func CreateDeployment(i IPlatform, build *data.BuildDefinition) (*beta.Deployment, error) {
	return i.CreateDeployment(build)
}

//CreateService within the platform cluster
func CreateService(i IPlatform, build *data.BuildDefinition) error {
	return i.CreateService(build)
}

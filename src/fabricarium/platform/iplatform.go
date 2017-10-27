package platform

import "github.com/AlexsJones/kubebuilder/src/data"

//IPlatform interface for container platform
type IPlatform interface {
	ValidateDeployment(build *data.BuildDefinition) (bool, error)
	CreateNamespace(string) error
	CreateDeployment(build *data.BuildDefinition) error
}

//ValidateDeployment from deserialisation of YAML
func ValidateDeployment(i IPlatform, build *data.BuildDefinition) (bool, error) {
	return i.ValidateDeployment(build)
}

//CreateNamespace within the platform cluster
func CreateNamespace(i IPlatform, namespace string) error {
	return i.CreateNamespace(namespace)
}

//CreateDeployment within the platform cluster
func CreateDeployment(i IPlatform, build *data.BuildDefinition) error {
	return i.CreateDeployment(build)
}

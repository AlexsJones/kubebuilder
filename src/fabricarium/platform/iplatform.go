package platform

import "github.com/AlexsJones/kubebuilder/src/data"

//IPlatform interface for container platform
type IPlatform interface {
	CreateNamespace(string) error
	CreateDeployment(build *data.BuildDefinition) error
}

//CreateNamespace within the platform cluster
func CreateNamespace(i IPlatform, namespace string) error {
	return i.CreateNamespace(namespace)
}

//CreateDeployment within the platform cluster
func CreateDeployment(i IPlatform, build *data.BuildDefinition) error {
	return i.CreateDeployment(build)
}

package vcs

import "github.com/AlexsJones/kubebuilder/src/data"

//IVCS ...
type IVCS interface {
	Fetch(checkoutPath string, build *data.BuildDefinition) error
}

//Fetch ...
func Fetch(iface IVCS, checkoutPath string, build *data.BuildDefinition) error {
	return iface.Fetch(checkoutPath, build)
}

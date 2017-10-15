package vcs

import (
	"os"
	"path"

	"github.com/AlexsJones/kubebuilder/src/data"
	git "gopkg.in/src-d/go-git.v4"
)

//GitVCS ...
type GitVCS struct {
}

//Fetch ...
func (g GitVCS) Fetch(checkoutPath string, build *data.BuildDefinition) error {
	p := path.Join(checkoutPath, build.VCS.Name)

	_, err := git.PlainClone(p, false, &git.CloneOptions{
		URL:      build.VCS.Path,
		Progress: os.Stdout,
	})

	return err
}

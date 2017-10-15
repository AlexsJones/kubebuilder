package vcs

import (
	"os"

	"github.com/AlexsJones/kubebuilder/src/data"
	git "gopkg.in/src-d/go-git.v4"
)

//GitVCS ...
type GitVCS struct {
}

//Fetch ...
func (g GitVCS) Fetch(checkoutPath string, build *data.BuildDefinition) error {

	_, err := git.PlainClone(checkoutPath, false, &git.CloneOptions{
		URL:          build.VCS.Path,
		Progress:     os.Stdout,
		SingleBranch: true,
	})

	return err
}

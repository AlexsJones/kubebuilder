package fabricarium_test

import (
	"github.com/AlexsJones/kubebuilder/src/fabricarium"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fabricarium", func() {
	It("should load correctly", func() {

		fab := fabricarium.NewFabricarium(nil)

		Expect(fab).NotTo(BeNil())
	})
})

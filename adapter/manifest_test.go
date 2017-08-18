package adapter_test

import (
	. "github.com/bstick12/go-patch-sdk/adapter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

var _ = Describe("Manifest", func() {

	var manifestGenerator ManifestGenerator

	BeforeEach(func() {
		manifestGenerator = ManifestGenerator{}
	})

	It("", func() {
		manifest, err := manifestGenerator.GenerateManifest(
			serviceadapter.ServiceDeployment{},
			serviceadapter.Plan{},
			serviceadapter.RequestParameters{},
			&bosh.BoshManifest{},
			&serviceadapter.Plan{},
		)
		Expect(err).ToNot(HaveOccurred())
		Expect(manifest).ShouldNot(BeNil())
	})

})

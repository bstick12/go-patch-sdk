package adapter_test

import (
	. "github.com/bstick12/go-patch-sdk/adapter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

var _ = Describe("Dashboardurl", func() {

	var dashboardUrlGenerator DashboardUrlGenerator

	BeforeEach(func() {
		dashboardUrlGenerator = DashboardUrlGenerator{}
	})

	It("when generate url", func() {
		dashboardUrl, err := dashboardUrlGenerator.DashboardUrl("", serviceadapter.Plan{}, bosh.BoshManifest{})
		Expect(err).ToNot(HaveOccurred())
		Expect(dashboardUrl).ToNot(BeNil())
	})

})

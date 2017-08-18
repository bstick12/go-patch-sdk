package adapter_test

import (
	. "github.com/bstick12/go-patch-sdk/adapter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
)

var _ = Describe("Binder", func() {

	It("when creating binding", func() {
		binder := Binder{}
		binding, err := binder.CreateBinding("", nil, bosh.BoshManifest{}, nil)
		Expect(err).NotTo(HaveOccurred())
		Expect(binding).NotTo(BeNil())
	})

	It("when creating binding", func() {
		binder := Binder{}
		err := binder.DeleteBinding("", nil, bosh.BoshManifest{}, nil)
		Expect(err).NotTo(HaveOccurred())
	})

})

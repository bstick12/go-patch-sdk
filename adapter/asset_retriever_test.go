package adapter_test

import (
	. "github.com/bstick12/go-patch-sdk/adapter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AssetRetriever", func() {

	Context("AssetRetriever retrieves files", func() {

		It("succeeds in retrieving file", func() {
			assertRetriever := NewAssertRetriever()
			_, err := assertRetriever.Asset("kubo.yml")
			Expect(err).NotTo(HaveOccurred())
		})

		It("fails to retrieve file", func() {
			assertRetriever := NewAssertRetriever()
			_, err := assertRetriever.Asset("not_found.yml")
			Expect(err).To(HaveOccurred())
		})

	})

})

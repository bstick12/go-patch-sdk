package adapter_test

import (
	. "github.com/bstick12/go-patch-sdk/adapter"
	"github.com/bstick12/go-patch-sdk/adapter/adapterfakes"

	"github.com/cloudfoundry/bosh-utils/errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

var _ = Describe("Manifest", func() {

	var manifestGenerator ManifestGenerator
	var fakeAssetRetriever = new(adapterfakes.FakeAssetRetriever)

	BeforeEach(func() {
		manifestGenerator = ManifestGenerator{}
		fakeAssetRetriever = new(adapterfakes.FakeAssetRetriever)
		manifestGenerator.AssetRetriever = fakeAssetRetriever
	})

	It("generates BOSH manifest", func() {

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

	It("raises error if kubo.yml not found", func() {

		fakeAssetRetriever.AssetReturnsOnCall(0, make([]byte, 0, 0),
			errors.Error("Failed to retrieve asset"))
		manifest, err := manifestGenerator.GenerateManifest(
			serviceadapter.ServiceDeployment{},
			serviceadapter.Plan{},
			serviceadapter.RequestParameters{},
			&bosh.BoshManifest{},
			&serviceadapter.Plan{},
		)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(Equal("Failed to load kubo.yml"))
		Expect(manifest).ShouldNot(BeNil())

	})

	It("raises error if kubo.yml is not valid yaml", func() {

		fakeAssetRetriever.AssetReturnsOnCall(0, []byte("This is not valid yaml."), nil)
		manifest, err := manifestGenerator.GenerateManifest(
			serviceadapter.ServiceDeployment{},
			serviceadapter.Plan{},
			serviceadapter.RequestParameters{},
			&bosh.BoshManifest{},
			&serviceadapter.Plan{},
		)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Failed to unmarshall"))
		Expect(manifest).ShouldNot(BeNil())

	})

	It("loads ops files", func() {
		manifest, err := manifestGenerator.GenerateManifest(
			serviceadapter.ServiceDeployment{},
			serviceadapter.Plan{},
			serviceadapter.RequestParameters{},
			&bosh.BoshManifest{},
			&serviceadapter.Plan{},
		)
		Expect(err).ToNot(HaveOccurred())
		assetName := fakeAssetRetriever.AssetArgsForCall(0)
		Expect(assetName).Should(Equal("kubo.yml"))
		assetName = fakeAssetRetriever.AssetArgsForCall(1)
		Expect(assetName).Should(Equal("ops-files/gcp-cloud-provider.yml"))
		assetName = fakeAssetRetriever.AssetArgsForCall(2)
		Expect(assetName).Should(Equal("ops-files/cf-routing.yml"))
		assetName = fakeAssetRetriever.AssetArgsForCall(3)
		Expect(assetName).Should(Equal("ops-files/remove-haproxy.yml"))
		Expect(manifest).ShouldNot(BeNil())
	})

	It("fails when asset doesn't exist", func() {
		fakeAssetRetriever.AssetReturnsOnCall(1, make([]byte, 0, 0),
			errors.Error("Failed to retrieve ops file"))
		manifest, err := manifestGenerator.GenerateManifest(
			serviceadapter.ServiceDeployment{},
			serviceadapter.Plan{},
			serviceadapter.RequestParameters{},
			&bosh.BoshManifest{},
			&serviceadapter.Plan{},
		)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Failed to load ops file"))
		Expect(manifest).ShouldNot(BeNil())

	})

	It("fails when ops file isn't valid", func() {
		fakeAssetRetriever.AssetReturnsOnCall(1, []byte("invalid ops"), nil)
		manifest, err := manifestGenerator.GenerateManifest(
			serviceadapter.ServiceDeployment{},
			serviceadapter.Plan{},
			serviceadapter.RequestParameters{},
			&bosh.BoshManifest{},
			&serviceadapter.Plan{},
		)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Failed to unmarshal ops file"))
		Expect(manifest).ShouldNot(BeNil())

	})

	It("fails when ops file can't be applied", func() {
		fakeAssetRetriever.AssetReturnsOnCall(1, []byte("- type: invalid"), nil)
		manifest, err := manifestGenerator.GenerateManifest(
			serviceadapter.ServiceDeployment{},
			serviceadapter.Plan{},
			serviceadapter.RequestParameters{},
			&bosh.BoshManifest{},
			&serviceadapter.Plan{},
		)
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("Failed to create ops from definitions"))
		Expect(manifest).ShouldNot(BeNil())

	})

})

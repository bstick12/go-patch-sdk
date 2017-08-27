package adapter

import (
	"log"

	"github.com/cloudfoundry/bosh-cli/director/template"
	"github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cppforlife/go-patch/patch"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
	"gopkg.in/yaml.v2"
)

type ManifestGenerator struct {
	StderrLogger   *log.Logger
	AssetRetriever AssetRetriever
}

func (mg ManifestGenerator) GenerateManifest(serviceDeployment serviceadapter.ServiceDeployment, plan serviceadapter.Plan,

	requestParams serviceadapter.RequestParameters, previousManifest *bosh.BoshManifest,
	previousPlan *serviceadapter.Plan) (bosh.BoshManifest, error) {

	kuboBytes, err := mg.AssetRetriever.Asset("kubo.yml")
	if err != nil {
		return bosh.BoshManifest{}, errors.Error("Failed to load kubo.yml")
	}

	ops, err := getOps(serviceDeployment, mg)
	if err != nil {
		return bosh.BoshManifest{}, err
	}

	return getManifest(plan.Properties, kuboBytes, ops)

}

func getManifest(properties serviceadapter.Properties,
	bytes []byte, ops patch.Ops) (bosh.BoshManifest, error) {

	tpl := template.NewTemplate(bytes)
	staticVariables := template.StaticVariables(properties)
	result, err := tpl.Evaluate(staticVariables, ops, template.EvaluateOpts{})
	if err != nil {
		return bosh.BoshManifest{}, errors.Errorf("Failed to apply ops files: %s", err.Error())
	}

	var boshManifest bosh.BoshManifest
	err = yaml.Unmarshal(result, &boshManifest)
	if err != nil {
		return bosh.BoshManifest{}, errors.Errorf("Failed to unmarshall %s due to %s ", string(bytes), err.Error())
	}
	return boshManifest, nil
}

func getOps(serviceDeployment serviceadapter.ServiceDeployment, manifestGenerator ManifestGenerator) (patch.Ops, error) {

	opsFiles := getOpsFiles(serviceDeployment)

	var op []patch.Op

	for _, opFile := range opsFiles {
		var opDefinitions []patch.OpDefinition
		opsBytes, err := manifestGenerator.AssetRetriever.Asset(opFile)
		if err != nil {
			return patch.Ops{}, errors.Errorf("Failed to load ops file %s due to %s", opFile, err.Error())
		}
		err = yaml.Unmarshal(opsBytes, &opDefinitions)
		if err != nil {
			return patch.Ops{}, errors.Errorf("Failed to unmarshal ops file %s due to %s", opFile, err.Error())
		}
		tmpOps, err := patch.NewOpsFromDefinitions(opDefinitions)
		if err != nil {
			return patch.Ops{}, errors.Errorf("Failed to create ops from definitions in %s due to %s", opFile, err.Error())
		}
		op = append(op, tmpOps)

	}

	return patch.Ops(op), nil

}

func getOpsFiles(serviceDeployment serviceadapter.ServiceDeployment) []string {
	return []string{
		"ops-files/gcp-cloud-provider.yml",
		"ops-files/cf-routing.yml",
		"ops-files/remove-haproxy.yml",
	}
}

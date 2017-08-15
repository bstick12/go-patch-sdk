package main

import (
	"fmt"

	"github.com/bstick12/go-patch-sdk/kubo"
	"github.com/cloudfoundry/bosh-cli/director/template"
	. "github.com/cppforlife/go-patch/patch"
	"github.com/pivotal-cf/on-demand-services-sdk/bosh"
	"gopkg.in/yaml.v2"
)

func main() {

	var opDefs []OpDefinition

	inBytes, err := kubo.Asset("kubo.yml")
	if err != nil {
		panic(err)
	}

	opsBytes, err := kubo.Asset("ops-files/gcp-cloud-provider.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(opsBytes, &opDefs)
	if err != nil {
		panic(err)
	}

	ops, err := NewOpsFromDefinitions(opDefs)
	if err != nil {
		panic(err)
	}

	var manifest interface{}

	err = yaml.Unmarshal(inBytes, &manifest)
	if err != nil {
		panic(err)
	}

	manifest, err = ops.Apply(manifest)
	if err != nil {
		panic(err)
	}

	str, err := yaml.Marshal(manifest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", str)

	var boshManifest bosh.BoshManifest
	err = yaml.Unmarshal([]byte(str), &boshManifest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", boshManifest)

	tpl := template.NewTemplate(inBytes)
	result, err := tpl.Evaluate(template.StaticVariables{}, ops, template.EvaluateOpts{
		ExpectAllKeys:     true,
		ExpectAllVarsUsed: true,
	})
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(result, &manifest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", manifest)

}

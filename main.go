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

	inBytes, err := kubo.Asset("kubo.yml")
	if err != nil {
		panic(err)
	}

	opsFiles := []string{
		// Not sure how we would pass the IAAS
		"ops-files/gcp-cloud-provider.yml",
		"ops-files/cf-routing.yml",
		"ops-files/remove-haproxy.yml",
	}

	var op []Op

	for _, opFile := range opsFiles {
		var opDefs []OpDefinition
		opsBytes, err := kubo.Asset(opFile)
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(opsBytes, &opDefs)
		if err != nil {
			panic(err)
		}

		tmpOps, err := NewOpsFromDefinitions(opDefs)
		if err != nil {
			panic(err)
		}
		op = append(op, tmpOps)

	}

	var ops = Ops(op)

	// These would come from the adapter.ServiceProperties
	staticVariables := template.StaticVariables{
		"deployment_name":              "deployment_name",
		"deployments_network":          "kubo-pcf",
		"kubernetes_master_port":       "8443",
		"network":                      "network",
		"project_id":                   "project_id",
		"stemcell_version":             "3412.11",
		"cf-tcp-router-name":           "cf-tcp-router-name",
		"routing-cf-api-url":           "routing-cf-api-url",
		"routing-cf-app-domain-name":   "routing-cf-app-domain-name",
		"routing-cf-client-id":         "routing-cf-client-id",
		"routing-cf-client-secret":     "routing-cf-client-secret",
		"routing-cf-nats-internal-ips": "routing-cf-nats-internal-ips",
		"routing-cf-nats-password":     "routing-cf-nats-password",
		"routing-cf-nats-port":         "routing-cf-nats-port",
		"routing-cf-nats-username":     "routing-cf-nats-username",
		"routing-cf-uaa-url":           "routing-cf-uaa-url",
	}

	tpl := template.NewTemplate(inBytes)
	result, err := tpl.Evaluate(
		staticVariables,
		ops,
		template.EvaluateOpts{})
	if err != nil {
		panic(err)
	}

	var boshManifest bosh.BoshManifest

	err = yaml.Unmarshal(result, &boshManifest)
	if err != nil {
		panic(err)
	}

	bosh, err := yaml.Marshal(boshManifest)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", string(bosh))
	fmt.Printf("--------------------------------------------\n")

}

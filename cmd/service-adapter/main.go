package main

import (
	"log"
	"os"

	"github.com/bstick12/go-patch-sdk/adapter"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

func main() {

	stderrLogger := log.New(os.Stderr, "[kubo-service-adapter]", log.LstdFlags)

	manifestGenerator := &adapter.ManifestGenerator{
		StderrLogger: stderrLogger,
		AssetRetriever: adapter.NewAssertRetriever(),
	}

	binder := &adapter.Binder{}

	serviceadapter.HandleCommandLineInvocation(
		os.Args,
		manifestGenerator,
		binder,
		&adapter.DashboardUrlGenerator{},
	)

}

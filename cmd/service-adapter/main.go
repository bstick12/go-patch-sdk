package main

import (
	"github.com/bstick12/go-patch-sdk/adapter"
	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
	"log"
	"os"
)

func main() {

	stderrLogger := log.New(os.Stderr, "[kubo-service-adapter]", log.LstdFlags)

	manifestGenerator := &adapter.ManifestGenerator{
		StderrLogger: stderrLogger,
	}

	binder := &adapter.Binder{}

	serviceadapter.HandleCommandLineInvocation(
		os.Args,
		manifestGenerator,
		binder,
		&adapter.DashboardUrlGenerator{},
	)

}

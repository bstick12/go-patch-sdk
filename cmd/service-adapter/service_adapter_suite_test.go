package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestServiceAdapter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ServiceAdapter Suite")
}

package local_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/sanksons/tavern/adapters/local"
)

var Adapter *Local

func TestLocal(t *testing.T) {
	InitializeAdapter()
	RegisterFailHandler(Fail)

	RunSpecs(t, "Local Suite")
}

func InitializeAdapter() {
	Adapter = Initialize(LocalAdapterConfig{})
}

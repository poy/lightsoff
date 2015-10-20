package lightsoff_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLightsoff(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lightsoff Suite")
}

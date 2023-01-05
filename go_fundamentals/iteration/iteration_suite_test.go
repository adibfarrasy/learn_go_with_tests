package iteration_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIteration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iteration Suite")
}

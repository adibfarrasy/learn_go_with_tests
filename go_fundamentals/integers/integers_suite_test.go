package integers_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestIntegers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Integers Suite")
}

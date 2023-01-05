package banking_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBanking(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Banking Suite")
}

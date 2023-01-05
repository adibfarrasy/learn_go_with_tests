package iteration_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"iteration/service"
)

var _ = Describe("Iteration feature", func() {
	When("user wants to repeat a word", func() {
		Context("given the word test and repetition of 5", func() {
			It("should repeat the word test 5 times", func() {
				Expect(service.Repeat("test", 5)).To(Equal("testtesttesttesttest"))
			})
		})
	})
})

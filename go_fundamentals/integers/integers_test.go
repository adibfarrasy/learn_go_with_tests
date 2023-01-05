package integers_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	a "integers/adder"
)

var _ = Describe("Integers feature", func() {
	When("user adds 2 + 2", func() {
		It("gets the result 4", func() {
			Expect(a.Add(2, 2)).To(Equal(4))
		})
	})
})

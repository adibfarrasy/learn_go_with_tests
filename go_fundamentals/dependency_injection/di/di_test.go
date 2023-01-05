package di_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"bytes"
	"dependency_injection/di"
)

var _ = Describe("Di", func() {
	When("user wants to greet", func() {
		Context("given Chris as name", func() {
			It("returns Hello, Chris", func() {
				buffer := bytes.Buffer{}
				di.Greet(&buffer, "Chris")

				Expect(buffer.String()).To(Equal("Hello, Chris"))
			})
		})
	})
})

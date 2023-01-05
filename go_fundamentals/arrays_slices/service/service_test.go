package service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"arrays-slices/service"
)

var _ = Describe("Arrays and slices feature", func() {
	When("a user wants to add numbers", func() {
		Context("given a list of ints", func() {
			It("gets the sum of the list of ints", func() {
				numbers := []int{1, 2, 3, 4, 5}
				Expect(service.Sum(numbers)).To(Equal(15))
			})
		})
	})

})

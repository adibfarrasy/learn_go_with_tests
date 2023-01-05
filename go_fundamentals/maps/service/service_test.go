package service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("[Dictionary Service]", func() {
	When("user wants to read from dictionary", func() {
		Context("given a dictionary of test: hello_world", func() {
			It("should return hello_world when test is called", func() {
				dictionary := map[string]string{"test": "hello_world"}
				Expect(dictionary["test"]).To(Equal("hello_world"))
			})
		})

		Context("given a dictionary of test: hello_world", func() {
			It("should return empty string when unknown_key is called", func() {
				dictionary := map[string]string{"test": "hello_world"}
				result, success := dictionary["unknown_key"]

				Expect(result).To(Equal(""))
				Expect(success).To(Equal(false))
			})
		})
	})

})

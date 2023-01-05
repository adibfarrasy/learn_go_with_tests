package main_test

import (
	hello "hello/service"
	data "hello/service/data"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hello feature", func() {
	When("user wants to say hello", func() {
		Context("and user supplied their name", func() {
			It("is successful with supplied name", func() {
				name := "test_name"
				Expect(hello.Hello(&name, nil)).To(Equal("Hello, test_name"))
			})
		})
		Context("and user doesn't supply their name", func() {
			It("is successful with default hello", func() {
				Expect(hello.Hello(nil, nil)).To(Equal("Hello, world"))
			})
		})
	})

	When("user wants to say hello in Spanish", func() {
		Context("and user supplied their name", func() {
			It("is successful with supplied name", func() {
				var name = "test_name"
				lang := data.Spanish
				Expect(hello.Hello(&name, &lang)).To(Equal("Hola, test_name"))
			})
		})
		Context("and user doesn't supply their name", func() {
			It("is successful with default hello", func() {
				name := ""
				lang := data.Spanish
				Expect(hello.Hello(&name, &lang)).To(Equal("Hola, mundo"))
			})
		})
	})
})

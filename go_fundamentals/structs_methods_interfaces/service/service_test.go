package service_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"structs_methods_interfaces/service"
)

var _ = Describe("Dimension service", func() {
	When("user wants to calculate perimeter", func() {
		Context("given a rectangle with width and height of 10 units", func() {
			It("should return perimeter of 40 units", func() {
				rectangle := service.Rectangle{10.0, 10.0}
				perimeter := service.RectanglePerimeter(rectangle)
				Expect(perimeter).To(Equal(40.0))
			})
		})
	})

	DescribeTable("user wants to calculate area",
		func(shape service.Shape, expected float64) {
			Expect(shape.Area()).To(Equal(expected))
		},
		Entry("given a rectangle with width and height of 10 units", service.Rectangle{10.0, 10.0}, 100.0),
		Entry("given a circle with radius of 10 units", service.Circle{10.0}, 314.0),
		Entry("given a triangle with base and height of 10 units", service.Triangle{10.0, 10.0}, 50.0),
	)
})

package mock_test

import (
	"bytes"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"mocking/mock"
)

var _ = Describe("Mock", func() {
	When("user wants to countdown", func() {
		Context("given a default countdown", func() {
			It("returns 3, 2, 1, Go!", func() {
				buffer := &bytes.Buffer{}
				spySleeper := &mock.SpySleeper{}

				mock.Countdown(buffer, spySleeper)
				expected := "3\n2\n1\nGo!\n"

				Expect(buffer.String()).To(Equal(expected))
				Expect(spySleeper.Calls).To(Equal(3))
			})
		})
	})
})

var _ = Describe("Meta testing the test rigs", func() {
	Context("given a default countdown", func() {
		Context("and check for order of operations", func() {
			It("returns 3, 2, 1, Go!", func() {
				spySleepPrinter := &mock.SpyCountdown{}

				mock.Countdown(spySleepPrinter, spySleepPrinter)
				expected := []string{
					"write",
					"sleep",
					"write",
					"sleep",
					"write",
					"sleep",
					"write",
				}

				Expect(spySleepPrinter.Calls).To(Equal(expected))
			})
		})

		Context("given adjustable duration for spy sleeper set to 5s", func() {
			It("should follow the sleep time of 5s", func() {
				sleepTime := 5 * time.Second

				spyTime := &mock.SpyTime{}
				sleeper := mock.ConfigurableSleeper{sleepTime, spyTime.Sleep}
				sleeper.Sleep()

				Expect(spyTime.DurationSlept).To(Equal(sleepTime))
			})
		})
	})
})

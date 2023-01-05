package main

import (
	"os"
	"time"

	"mocking/mock"
)

func main() {
	sleeper := &mock.ConfigurableSleeper{Duration: time.Second, SleepFn: time.Sleep}
	mock.Countdown(os.Stdout, sleeper)
}

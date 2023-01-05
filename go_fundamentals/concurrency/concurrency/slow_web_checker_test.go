package concurrency_test

import (
	"concurrency/concurrency"
	"testing"
	"time"
)

func SlowStubWebsiteChecker(_ string) bool {
	time.Sleep(2 * time.Millisecond)
	return true
}

func BenchmarkCheckWBenchmarkCheckW(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency.CheckWebsites(SlowStubWebsiteChecker, urls)
	}
}

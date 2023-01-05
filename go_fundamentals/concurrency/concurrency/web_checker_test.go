package concurrency_test

import (
	"concurrency/concurrency"
	"reflect"
	"testing"
)

func mockWebisteChecker(url string) bool {
	if url == "waat://" {
		return false
	}
	return true
}

func TestWebsiteChecker(t *testing.T) {
	websites := []string{
		"http://www.google.com",
		"http://www.test.cok",
		"waat://",
	}

	want := map[string]bool{
		"http://www.google.com": true,
		"http://www.test.cok":   true,
		"waat://":               false,
	}

	got := concurrency.CheckWebsites(mockWebisteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

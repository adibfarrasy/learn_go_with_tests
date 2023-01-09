package main

import (
	"testing"
	"time"

	"github.com/quii/go-graceful-shutdown/acceptancetests"
	"github.com/quii/go-graceful-shutdown/assert"
)

const (
	port = "8080"
	url  = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)

	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	// check the server works before shutting down
	assert.CanGet(t, url)

	// fire off a request, shutdown before responding
	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})

	// without graceful shutdown this will fail
	assert.CanGet(t, url)

	// no more request should work after sendInterrupt
	assert.CantGet(t, url)
}

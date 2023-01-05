package counter_test

import (
	"sync"
	. "sync/counter"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

        assertCounter(t, counter.Value(), 3)
	})

    t.Run("it runs safely concurrently", func(t *testing.T){
        want := 1000

        counter := Counter{}

        var wg sync.WaitGroup
        wg.Add(want)

        for i := 0; i < want; i++ {
            go func(){
                counter.Inc()
                wg.Done()
            }()
        }

        wg.Wait()

        assertCounter(t, counter.Value(), want)
    })
}

func assertCounter(t testing.TB, got int, want int) {
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
}

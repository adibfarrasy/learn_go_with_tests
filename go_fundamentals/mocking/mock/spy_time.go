package mock

import "time"

type SpyTime struct {
	DurationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.DurationSlept = duration
}

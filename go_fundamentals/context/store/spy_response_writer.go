package store

import (
	"errors"
	"net/http"
)

type SpyResponseWriter struct {
	Written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.Written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.Written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.Written = true
}

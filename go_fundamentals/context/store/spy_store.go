package store

import (
	"context"
	"log"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

type SpyStore struct {
	Response string
	T        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
    data := make(chan string, 1)

    go func(){
        var result string
        for _, c := range s.Response {
            select {
            case <- ctx.Done():
                log.Println("spy store got cancelled")
                return
            default:
                time.Sleep(10 * time.Millisecond)
                result += string(c)
            }
        }

        data <- result
    }()

    select {
    case <- ctx.Done():
        return "", ctx.Err()
    case res := <- data:
        return res, nil
    }
}

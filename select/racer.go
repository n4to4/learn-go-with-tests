package racer

import (
	"net/http"
	"time"
)

func Racer(a, b string) string {
	aDuration := measureResponse(a)
	bDuration := measureResponse(b)

	if aDuration < bDuration {
		return a
	}
	return b
}

func measureResponse(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

package httpcheck

import (
	"net/http"
	"time"
)

// GetResponseAndTime Get GetResponseData And GetResponseTime
func GetResponseAndTime(target string) (*http.Response, float64, error) {
	start := time.Now()
	resp, err := http.Get(target)
	if err != nil {
		return nil, 0, err
	}
	elapsed := time.Since(start).Seconds()
	return resp, elapsed, nil
}

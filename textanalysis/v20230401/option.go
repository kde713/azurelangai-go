package v20230401

import (
	"time"
)

type options struct {
	// Retry
	retryCount       int
	retryWaitTime    time.Duration
	retryMaxWaitTime time.Duration
}

type Option func(*options)

func WithRetryCount(count int, minWait time.Duration, maxWait time.Duration) Option {
	return func(o *options) {
		o.retryCount = count
		o.retryWaitTime = minWait
		o.retryMaxWaitTime = maxWait
	}
}

package logicmonitor

import (
	"net/http"
	"time"
)

// RetryTransportConfig configures the provider's HTTP 429 retry transport.
// Tests in package test use this to build the same transport the provider
// installs, without depending on unexported types in provider.go.
type RetryTransportConfig struct {
	// Underlying is the next RoundTripper in the chain.
	Underlying http.RoundTripper
	// MaxRetries is how many times a 429 may be retried after the first attempt.
	MaxRetries int
	// BaseDelay is the exponential-backoff starting delay when Retry-After is absent.
	BaseDelay time.Duration
	// MaxDelay caps every computed wait, including a Retry-After value.
	MaxDelay time.Duration
}

// NewRetryTransport returns the provider's 429 retry transport.
func NewRetryTransport(cfg RetryTransportConfig) http.RoundTripper {
	return &retryTransport{
		underlyingTransport: cfg.Underlying,
		maxRetries:          cfg.MaxRetries,
		baseDelay:           cfg.BaseDelay,
		maxDelay:            cfg.MaxDelay,
	}
}

// NewUserAgentTransport returns the provider's User-Agent transport.
func NewUserAgentTransport(underlying http.RoundTripper, userAgent string) http.RoundTripper {
	return &userAgentTransport{
		underlyingTransport: underlying,
		userAgent:           userAgent,
	}
}

// RetryBackoff returns the delay the retry transport waits before the given attempt.
// retryAfter is the raw Retry-After header value, or empty to use exponential backoff.
func RetryBackoff(baseDelay, maxDelay time.Duration, attempt int, retryAfter string) time.Duration {
	tr := &retryTransport{baseDelay: baseDelay, maxDelay: maxDelay}
	return tr.backoff(attempt, retryAfter)
}

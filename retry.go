package goretry

import "time"

type retryConfig struct {
	attempts int
	backoff  BackOff
	retryIf  RetryIf
	onRetry  OnRetry
}

func newDefaultRetryConfig() *retryConfig {
	return &retryConfig{
		attempts: 10,
		backoff:  nil,
		retryIf:  nil,
		onRetry:  nil,
	}
}

// goretry function
type Retry func(retries int) error

// default always goretry, and do nothing when error return
func Do(retryFunc Retry, options ...Option) error {
	var err error
	config := newDefaultRetryConfig()

	// option
	for _, opt := range options {
		opt(config)
	}

	for retries := 0; retries < config.attempts; retries++ {
		err = retryFunc(retries)

		if err != nil {
			if config.onRetry != nil {
				config.onRetry(retries, err)
			}

			if config.retryIf != nil && !config.retryIf(retries, err) {
				break
			}

			if config.backoff != nil {
				delay := config.backoff(retries)
				if delay > 0 {
					time.Sleep(delay)
				}
			}

		} else {
			return nil
		}
	}

	return err
}

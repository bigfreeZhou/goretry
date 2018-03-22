package goretry

type Option func(*retryConfig)

// RetryIf to determine if need retry next, return true, else not
type RetryIf func(retries int, err error) bool

type OnRetry func(retries int, err error)

func MaxRetryOption(max int) Option {
	return func(conf *retryConfig) {
		conf.attempts = max
	}
}

func BackOffOption(backoff BackOff) Option {
	return func(conf *retryConfig) {
		conf.backoff = backoff
	}
}

func OnRetryOption(onretry OnRetry) Option {
	return func(conf *retryConfig) {
		conf.onRetry = onretry
	}
}

func RetryIfOption(retryIf RetryIf) Option {
	return func(conf *retryConfig) {
		conf.retryIf = retryIf
	}
}

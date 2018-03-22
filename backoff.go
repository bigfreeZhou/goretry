package goretry

import (
	"math"
	"time"
)

// backoff returns the amount of time to wait before the next goretry given the number of consecutive failures.
type BackOff func(retries int) time.Duration

func NoStopBackOff() BackOff {
	return func(retries int) time.Duration {
		return 0
	}
}

func ConstantBackOff(delay time.Duration) BackOff {
	return func(retries int) time.Duration {
		return delay
	}
}

func IncrementBackOff(initial, increment time.Duration) BackOff {
	return func(retries int) time.Duration {
		return initial + (increment + time.Duration(retries))
	}
}

// linear increase delay duration
func LinearBackOff(factor time.Duration) BackOff {
	return func(retries int) time.Duration {
		return time.Duration(retries) * factor
	}
}

// exponential increase delay duration
func ExpBackOff(factor time.Duration, base float64) BackOff {
	return func(retries int) time.Duration {
		return (factor * time.Duration(math.Pow(base, float64(retries))))
	}
}

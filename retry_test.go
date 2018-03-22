package goretry

import (
	"fmt"
	"testing"
)

func TestSuccessOnce(t *testing.T) {
	err := Do(func(retries int) error {
		t.Logf("do goretry: %d", retries)
		return nil
	})

	t.Logf("after goretry, err: %v", err)
}

func TestSimpleRetry(t *testing.T) {
	err := Do(func(retries int) error {
		t.Logf("do goretry: %d", retries)
		return fmt.Errorf("test simple goretry error, retries: %d", retries)
	})

	t.Logf("error happen when do goretry, err: %v", err)
}

func TestMaxRetryOption(t *testing.T) {
	err := Do(func(retries int) error {
		t.Logf("do goretry: %d", retries)
		return fmt.Errorf("test simple goretry error, retries: %d", retries)
	}, MaxRetryOption(5))

	t.Logf("error happen when do goretry, err: %v", err)
}

func TestMaxRetryIfOption(t *testing.T) {
	err := Do(func(retries int) error {
		t.Logf("do goretry: %d", retries)
		return fmt.Errorf("test simple goretry error, retries: %d", retries)
	},
		MaxRetryOption(5),
		RetryIfOption(func(retries int, err error) bool {
			t.Logf("error print in on goretry if, err: %v", err)
			if retries < 4 {
				return true
			}
			return false
		}),
		OnRetryOption(func(retryIf int, err error) {
			t.Logf("error print in on goretry, err: %v", err)
		}))

	t.Logf("error happen when do goretry, err: %v", err)
}

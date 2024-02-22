package retry_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/missingstudio/studio/common/resilience"
	"github.com/missingstudio/studio/common/resilience/retry"
	"github.com/zeebo/assert"
)

var ErrRetryExecution = errors.New("wanted error")

type retryFailer struct {
	notFailOnAttemp int
	timesExecuted   int
}

func (c *retryFailer) Run(ctx context.Context) error {
	c.timesExecuted++
	if c.timesExecuted == c.notFailOnAttemp {
		return nil
	}

	return ErrRetryExecution
}

func TestRetryResult(t *testing.T) {
	tests := []struct {
		name      string
		cfg       retry.Config
		getF      func() resilience.Func
		expResult string
		expErr    error
	}{
		{
			name: "A failing execution should not fail if it's retried the required times until returns a non error.",
			cfg: retry.Config{
				WaitBase:       1 * time.Nanosecond,
				DisableBackoff: true,
				Numbers:        3,
			},
			getF: func() resilience.Func {
				c := &retryFailer{notFailOnAttemp: 4}
				return c.Run
			},
			expErr: nil,
		},
		{
			name: "A failing execution should fail if it's not retried the required times until returns a non error.",
			cfg: retry.Config{
				WaitBase:       1 * time.Nanosecond,
				DisableBackoff: true,
				Numbers:        3,
			},
			getF: func() resilience.Func {
				c := &retryFailer{notFailOnAttemp: 5}
				return c.Run
			},
			expErr: ErrRetryExecution,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := retry.New(test.cfg)
			err := cmd.Run(context.TODO(), test.getF())

			assert.True(t, errors.Is(test.expErr, err))
		})
	}
}

package resilience

import (
	"context"
	"errors"
)

var ErrContextCanceled = errors.New("context canceled, logic not executed")

type Func func(ctx context.Context) error

type Runner interface {
	Run(ctx context.Context, f Func) error
}

type (
	RunnerMiddleware func(Runner) Runner
	RunnerFunc       func(ctx context.Context, f Func) error
)

func (r RunnerFunc) Run(ctx context.Context, f Func) error {
	return r(ctx, f)
}

// command is the unit of execution.
type command struct{}

// Run satisfies Runner interface.
func (command) Run(ctx context.Context, f Func) error {
	// Only execute if we reached to the execution and the context has not been cancelled.
	select {
	case <-ctx.Done():
		return ErrContextCanceled
	default:
		return f(ctx)
	}
}

// SanitizeRunner returns a safe execution Runner if the runner is nil.
// Usually this helper will be used for the last part of the runner chain
// when the runner is nil, so instead of acting on a nil Runner its executed
// on a `command` Runner, this runner knows how to execute the `Func` function.
// It's safe to use it always as if it encounters a safe Runner it will return
// that Runner.
func SanitizeRunner(r Runner) Runner {
	// In case of end of execution chain.
	if r == nil {
		return &command{}
	}
	return r
}

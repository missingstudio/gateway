package interceptor

import (
	"fmt"

	"github.com/missingstudio/studio/common/errors"
)

var (
	ErrProviderHeaderNotExit = errors.New(fmt.Errorf("x-ms-provider provider header not available"))
	ErrRateLimitExceeded     = errors.NewForbidden("rate limit exceeded")
	ErrUnauthenticated       = errors.NewUnauthorized("unauthenticated")
)

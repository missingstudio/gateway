package errors

import (
	"fmt"

	"github.com/missingstudio/studio/backend/internal/constants"
	"github.com/missingstudio/studio/common/errors"
)

var (
	ErrProviderHeaderNotExit = errors.NewBadRequest(fmt.Sprintf("%s header is required", constants.XMSProvider))
	ErrRequiredHeaderNotExit = errors.NewBadRequest(fmt.Sprintf("either %s or %s header is required", constants.XMSProvider, constants.XMSConfig))
	ErrRateLimitExceeded     = errors.NewForbidden("rate limit exceeded")
	ErrUnauthenticated       = errors.NewUnauthorized("unauthenticated")
	ErrProviderNotFound      = errors.NewNotFound("provider is not found")
)

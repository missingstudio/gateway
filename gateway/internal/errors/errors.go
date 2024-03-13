package errors

import (
	"fmt"

	"github.com/missingstudio/ai/gateway/internal/constants"
	"github.com/missingstudio/common/errors"
)

var (
	ErrProviderHeaderNotExit = errors.NewBadRequest(fmt.Sprintf("%s header is required", constants.XMSProvider))
	ErrRequiredHeaderNotExit = errors.NewBadRequest(fmt.Sprintf("either %s or %s header is required", constants.XMSProvider, constants.XMSConfig))
	ErrRateLimitExceeded     = errors.NewForbidden("rate limit exceeded")
	ErrUnauthenticated       = errors.NewUnauthorized("you are not authorized to access APIs")
	ErrProviderNotFound      = errors.NewNotFound("provider is not found")
	ErrRouterConfigNotValid  = errors.NewNotFound("router config is not valid")
)

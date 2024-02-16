package providers

import (
	"net/http"

	"github.com/missingstudio/studio/backend/internal/providers/base"
)

type ProviderFactory interface {
	Create(headers http.Header) (base.ProviderInterface, error)
}

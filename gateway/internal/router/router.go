package router

import (
	"errors"

	"github.com/missingstudio/studio/backend/internal/providers/base"
)

var ErrNoHealthyProviders = errors.New("no healthy providers found")

type Strategy string

type ProviderIterator interface {
	Next() (base.IProvider, error)
}

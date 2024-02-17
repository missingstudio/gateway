package models

import (
	"github.com/missingstudio/studio/backend/internal/providers/base"
)

// ProviderRegistry holds all supported provider for which connections
// can be initialized
var ProviderRegistry = map[string]func(Connection) base.IProvider{}

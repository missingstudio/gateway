package v1

import "github.com/missingstudio/studio/backend/internal/ingester"

type Deps struct {
	ingester ingester.Ingester
}

func NewDeps(ingester ingester.Ingester) *Deps {
	return &Deps{
		ingester: ingester,
	}
}

package mock

import "github.com/missingstudio/studio/backend/internal/providers/base"

type ProviderMock struct {
	Name string
}

func NewProviderMock(name string) base.ProviderInterface {
	return &ProviderMock{
		Name: name,
	}
}

func (p ProviderMock) GetName() string {
	return p.Name
}

func (p ProviderMock) Validate() error {
	return nil
}

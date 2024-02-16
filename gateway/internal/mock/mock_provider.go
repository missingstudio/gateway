package mock

import "github.com/missingstudio/studio/backend/internal/providers/base"

var _ base.ProviderInterface = &providerMock{}

type providerMock struct {
	Name string
}

func NewProviderMock(name string) base.ProviderInterface {
	return &providerMock{
		Name: name,
	}
}

func (p providerMock) GetName() string {
	return p.Name
}

func (p providerMock) Validate() error {
	return nil
}

func (*providerMock) GetModels() []interface{} {
	return []interface{}{}
}

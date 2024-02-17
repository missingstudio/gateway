package mock

import "github.com/missingstudio/studio/backend/internal/providers/base"

var _ base.IProvider = &providerMock{}

type providerMock struct {
	name string
}

func NewProviderMock(name string) base.IProvider {
	return &providerMock{
		name: name,
	}
}

func (p providerMock) Name() string {
	return p.name
}

func (togetherAI providerMock) Schema() []byte {
	return []byte{}
}

func (p providerMock) Validate() error {
	return nil
}

func (*providerMock) Models() []string {
	return []string{}
}

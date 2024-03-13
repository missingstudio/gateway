package mock

import "github.com/missingstudio/ai/gateway/internal/provider/base"

var _ base.Provider = &providerMock{}

type providerMock struct {
	info   base.Info
	config base.Config
}

func NewProviderMock(name string) base.Provider {
	return &providerMock{
		info: base.Info{Name: name},
	}
}

func (p providerMock) Info() base.Info {
	return p.info
}

func (p providerMock) Config() base.Config {
	return p.config
}

func (p providerMock) Schema() []byte {
	return []byte{}
}

func (p providerMock) Models() []string {
	return []string{}
}

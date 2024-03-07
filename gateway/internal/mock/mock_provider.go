package mock

import "github.com/missingstudio/ai/gateway/internal/providers/base"

var _ base.IProvider = &providerMock{}

type providerMock struct {
	info   base.ProviderInfo
	config base.ProviderConfig
}

func NewProviderMock(name string) base.IProvider {
	return &providerMock{
		info: base.ProviderInfo{Name: name},
	}
}

func (p providerMock) Info() base.ProviderInfo {
	return p.info
}

func (p providerMock) Config() base.ProviderConfig {
	return p.config
}

func (p providerMock) Schema() []byte {
	return []byte{}
}

func (p providerMock) Models() []string {
	return []string{}
}

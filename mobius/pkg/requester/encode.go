package requester

import (
	"encoding/json"
)

type (
	Encoder func(any) ([]byte, error)
	Decoder func([]byte, any) error
)

var (
	defaultEncoder = json.Marshal
	defaultDecoder = json.Unmarshal
)

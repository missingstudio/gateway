package models

import (
	"encoding/json"

	"github.com/Jeffail/gabs/v2"
	"github.com/google/uuid"
)

const (
	AuthorizationHeader string = "Authorization"
)

type ConnectionState string

const (
	ConnectionStateActive   ConnectionState = "active"
	ConnectionStateInactive ConnectionState = "inactive"
)

func (s ConnectionState) String() string {
	return string(s)
}

type Connection struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`

	// Config stores user input configurations
	Headers map[string]any  `json:"headers"`
	Config  map[string]any  `json:"config"`
	State   ConnectionState `json:"state"`
}

func (c *Connection) MergeConfig(kv map[string]any) (err error) {
	container := gabs.New()
	for k, v := range kv {
		_, err = container.SetP(v, k)
		if err != nil {
			return err
		}
	}
	if err := container.MergeFn(gabs.Wrap(c.Config), func(destination, source any) any {
		return destination
	}); err != nil {
		return err
	}

	// get back connection config
	containerBytes, err := container.MarshalJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal(containerBytes, &c.Config)
}

func (c *Connection) GetHeaders(keys []string) map[string]any {
	fetched := map[string]any{}
	container := gabs.Wrap(c.Headers)
	for _, key := range keys {
		fetched[key] = container.Path(key).Data()
	}
	return fetched
}

func (c *Connection) GetConfig(keys []string) map[string]any {
	fetched := map[string]any{}
	container := gabs.Wrap(c.Config)
	for _, key := range keys {
		fetched[key] = container.Path(key).Data()
	}
	return fetched
}

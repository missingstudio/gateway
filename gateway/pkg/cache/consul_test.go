package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/consul/api"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewConsulClient(t *testing.T) {
	config := &ConsulConfig{}

	client, err := NewConsulClient(config)
	assert.NotNil(t, client)
	assert.Nil(t, err)
}

func TestSetSuccess(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/kv/k1", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "true")
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	config := ConsulConfig{
		Config: api.Config{
			Address: ts.URL,
		},
	}
	c, err := NewConsulClient(&config)
	assert.NotNil(t, c)
	assert.Nil(t, err)

	ctx := context.Background()
	err = c.Set(ctx, "k1", []byte("v1"), 0)
	assert.Nil(t, err)
}

func TestSetFailure(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/kv/k1", func(res http.ResponseWriter, req *http.Request) {
		http.Error(res, "request failed", http.StatusInternalServerError)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	config := ConsulConfig{
		Config: api.Config{
			Address: ts.URL,
		},
	}
	c, err := NewConsulClient(&config)
	assert.NotNil(t, c)
	assert.Nil(t, err)

	ctx := context.Background()
	err = c.Set(ctx, "k1", []byte("v1"), 0)
	assert.NotNil(t, err)
}

func TestGetSuccess(t *testing.T) {
	expectedVal := []byte("v1")

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/kv/k1", func(res http.ResponseWriter, req *http.Request) {
		response := api.KVPairs{
			{
				Key:   "k1",
				Value: expectedVal,
			},
		}

		resBytes, err := json.Marshal(response)
		assert.Nil(t, err)
		res.Write(resBytes)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	config := ConsulConfig{
		Config: api.Config{
			Address: ts.URL,
		},
	}
	c, err := NewConsulClient(&config)
	assert.NotNil(t, c)
	assert.Nil(t, err)

	ctx := context.Background()
	val, err := c.Get(ctx, "k1")
	t.Log(val)
	assert.Equal(t, expectedVal, val)
	assert.Nil(t, err)
}

func TestGetFailure(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/kv/k1", func(res http.ResponseWriter, req *http.Request) {
		http.Error(res, "request failed", http.StatusInternalServerError)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	config := ConsulConfig{
		Config: api.Config{
			Address: ts.URL,
		},
	}
	c, err := NewConsulClient(&config)
	assert.NotNil(t, c)
	assert.Nil(t, err)

	ctx := context.Background()
	val, err := c.Get(ctx, "k1")
	t.Log(val)
	assert.Nil(t, val)
	assert.NotNil(t, err)
}

func TestDeleteSuccess(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/kv/k1", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "true")
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	config := ConsulConfig{
		Config: api.Config{
			Address: ts.URL,
		},
	}
	c, err := NewConsulClient(&config)
	assert.NotNil(t, c)
	assert.Nil(t, err)

	ctx := context.Background()
	err = c.Delete(ctx, "k1")
	assert.Nil(t, err)
}

func TestDeleteFailure(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/kv/k1", func(res http.ResponseWriter, req *http.Request) {
		http.Error(res, "request failed", http.StatusInternalServerError)
	})

	ts := httptest.NewServer(mux)
	defer ts.Close()

	config := ConsulConfig{
		Config: api.Config{
			Address: ts.URL,
		},
	}
	c, err := NewConsulClient(&config)
	assert.NotNil(t, c)
	assert.Nil(t, err)

	ctx := context.Background()
	err = c.Delete(ctx, "k1")
	assert.NotNil(t, err)
}

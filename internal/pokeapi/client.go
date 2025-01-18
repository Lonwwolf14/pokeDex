package pokeapi

import (
	"net/http"
	"time"
)

type Client struct {
	cache      *Cache
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

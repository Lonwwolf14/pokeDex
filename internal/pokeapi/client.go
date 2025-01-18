package pokeapi

import (
	"net/http"
	"time"

	pokecache "example.com/pokedex/internal/pokecache"
)

type Client struct {
	cache      *pokecache.Cache
	httpClient *http.Client
}

func NewClient(timeout, cacheInterval time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{},
		cache:      pokecache.NewCache(cacheInterval),
	}
}

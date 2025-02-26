package pokeapi

import (
	"net/http"
	"time"

	"github.com/joeljosephwebdev/pokedex-repl/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeOut, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeOut,
		},
	}
}

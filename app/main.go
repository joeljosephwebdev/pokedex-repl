package main

import (
	"time"

	"github.com/joeljosephwebdev/pokedex-repl/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &Config{
		pokeapiClient: pokeClient,
	}

	start_repl(cfg)
}

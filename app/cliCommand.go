package main

import "github.com/joeljosephwebdev/pokedex-repl/internal/pokeapi"

// cliCommand provides the core command registration and management system
// for the Pokedex application.

// cliCommand represents a single command in the Pokedex CLI.
// It contains the command's name, description, and the function to execute.
type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

func (c cliCommand) Name() string {
	return c.name
}

func (c cliCommand) Description() string {
	return c.description
}

func (c cliCommand) Callback() func(*Config, []string) error {
	return c.callback
}

var CommandsList = map[string]cliCommand{}

// Register adds a new command to the CommandsList map.
// It takes a name, description, and callback function as parameters.
func Register(name string, description string, callback func(*Config, []string) error) {
	CommandsList[name] = cliCommand{
		name:        name,
		description: description,
		callback:    callback,
	}
}

// Config maintains state for paginated API responses.
type Config struct {
	pokeapiClient        pokeapi.Client
	nextLocationsURL     *string
	previousLocationsURL *string
}

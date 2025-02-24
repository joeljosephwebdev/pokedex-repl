module github.com/joeljosephwebdev/pokedex-repl

go 1.23.5

replace (
	github.com/joeljosephwebdev/pokedex-repl/cleanInput v0.0.0 => ./internal/cleanInput
	github.com/joeljosephwebdev/pokedex-repl/internal/commands v0.0.0 => ./internal/commands
)

package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joeljosephwebdev/pokedex-repl/internal/cleanInput"
	"github.com/joeljosephwebdev/pokedex-repl/internal/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		cleaned := cleanInput.CleanInput(scanner.Text())
		if len(cleaned) < 1 {
			fmt.Println("Please enter a command. For a list of commands type 'help'.")
		} else {
			if command, exists := commands.Commands[cleaned[0]]; exists {
				callback := command.Callback()
				err := callback(cleaned[1:])
				if err != nil {
					fmt.Printf("Failed to execute command - %v\n", err)
				}
			} else {
				fmt.Printf("Invalid command %s. For a list of commands type 'help'.\n", cleaned[0])
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}

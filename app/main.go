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
	config := commands.Config{
		Next:     nil,
		Previous: nil,
	}

	//start repl loop
	for {
		fmt.Print("Pokedex > ")

		//wait for user input
		scanner.Scan()
		cleaned := cleanInput.CleanInput(scanner.Text())

		//check if input isn't empty
		if len(cleaned) < 1 {
			fmt.Println("Please enter a command. For a list of commands type 'help'.")
		} else {
			//check if input contains a valid command
			if command, exists := commands.Commands[cleaned[0]]; exists {
				callback := command.Callback()
				err := callback(&config, cleaned[1:])
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

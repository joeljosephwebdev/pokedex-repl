package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func start_repl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	//start repl loop
	for {
		fmt.Print("Pokedex > ")
		//wait for user input
		scanner.Scan()
		words := cleanInput(scanner.Text())
		//check if input isn't empty
		if len(words) < 1 {
			continue
		}
		//check if input contains a valid command
		commandName := words[0]
		if command, exists := CommandsList[commandName]; exists {
			callback := command.Callback()
			err := callback(cfg, words[1:])
			if err != nil {
				fmt.Printf("Failed to execute command - %v\n", err)
			}
		} else {
			fmt.Printf("Invalid command %s. For a list of commands type 'help'.\n", words[0])
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}

func cleanInput(text string) []string {
	lower := strings.ToLower((text))
	words := strings.Fields(lower)

	return words
}

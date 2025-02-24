package cleanInput

import (
	"strings"
)

func CleanInput(text string) []string {
	lower := strings.ToLower((text))
	words := strings.Fields(lower)

	return words
}

package cleanInput

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "toomanywordswithNOspacesinbetweenANDSOMECHARSARECapitaLIZED",
			expected: []string{"toomanywordswithnospacesinbetweenandsomecharsarecapitalized"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		// Check the length of the actual slice
		if len(actual) != len(c.expected) {
			// if they don't match, use t.Errorf to print an error message
			// fmt.Printf("%t - %t", actual, c.expected)
			t.Errorf("The length of the return value, %d, does not match the expected length %d.", len(actual), len(c.expected))
			return
		}
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			if word != expectedWord {
				// if they don't match, use t.Errorf to print an error message
				t.Errorf("Word at index %d, %s, does not match expected word, %s,", i, word, expectedWord)
				// and fail the test
				return
			}
		}
	}
}

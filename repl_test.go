package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "  hello   world\n",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  pikachu is a pokemon",
			expected: []string{"pikachu", "is", "a", "pokemon"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len actual '%v' != len expected '%v'", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word '%v' did not match expected word '%v'", word, expectedWord)
			}
		}
	}
}

package main

import "testing"

func TestCleanInput(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{{
		input:    "   Hello Pokemon ",
		expected: []string{"hello", "pokemon"},
	}, {
		input:    "charizzard Ok pikACHu            ",
		expected: []string{"charizzard", "ok", "pikachu"},
	}, {
		input:    "charmander                  MewTwo ",
		expected: []string{"charmander", "mewtwo"},
	}}

	for _, c := range tests {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("got %v, want %v", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("got %v, want %v", word, expectedWord)
			}
		}
	}

}

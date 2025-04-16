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
            input:    "  hello  world  ",
            expected: []string{"hello", "world"},
        },
        {
            input:    "Charmander Bulbasaur PIKACHU",
            expected: []string{"charmander", "bulbasaur", "pikachu"},
        },
        {
            input:    "   OneWord   ",
            expected: []string{"oneword"},
        },
        {
            input:    "UPPER lower MiXeD",
            expected: []string{"upper", "lower", "mixed"},
        },
        {
            input:    "",
            expected: []string{},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)

        if len(actual) != len(c.expected) {
            t.Errorf("For input %q, expected length %d, got %d", c.input, len(c.expected), len(actual))
            continue
        }

        for i := range actual {
            if actual[i] != c.expected[i] {
                t.Errorf("For input %q, expected word %q at index %d, got %q", c.input, c.expected[i], i, actual[i])
            }
        }
    }
}
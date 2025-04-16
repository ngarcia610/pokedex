package main

import (
    "strings"
)

func cleanInput(text string) []string {
    trimmed := strings.TrimSpace(text)
    lower := strings.ToLower(trimmed)
    words := strings.Fields(lower)
    return words
}
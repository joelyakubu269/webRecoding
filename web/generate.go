package main

import (
	"fmt"
	"strings"
)

func Generate(s string, banner map[rune][]string) (string, error) {
	var res strings.Builder
	text := strings.ReplaceAll(s, "\r\n", "\n")
	words := strings.Split(text, "\n")

	for i, r := range words {
		if r == "" {
			if i < len(words)-1 {
				res.WriteString("\n")
			}
		} else {
			n, err := Render(s, banner)
			if err != nil {
				return "", err
			}
			res.WriteString(n)
		}
	}
	return res.String(), nil
}
func Render(s string, banner map[rune][]string) (string, error) {
	var res strings.Builder

	for i := 0; i < 8; i++ {
		for _, ch := range s {
			val, ok := banner[ch]
			if ok {
				res.WriteString(val[i])
			} else {
				return "", fmt.Errorf("character not found")
			}

		}
		res.WriteString("\n")
	}
	return res.String(), nil
}
func validate(s string) error {
	for _, r := range s {
		if r == '\r' || r == '\n' {
			continue
		}
		if r < 32 || r > 126 {
			return fmt.Errorf("invalid character")
		}
	}
	return nil
}

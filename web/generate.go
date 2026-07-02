package main

import (
	"fmt"
	"strings"
)

func Generate(s string, banner map[rune][]string) (string, error) {

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

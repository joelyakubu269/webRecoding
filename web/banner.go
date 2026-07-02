package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBnner(filename string) (map[rune][]string, error) {
	banner := "banners/" + filename + ".txt"
	data, err := os.ReadFile(banner)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) < 855 {
		return nil, fmt.Errorf("invalid banner format")
	}
	m := make(map[rune][]string)
	for i := 32; i <= 126; i++ {
		start := ((i - 32) * 9) + 1
		end := start + 8
		m[rune(i)] = lines[start:end]
	}
	return m, nil
}

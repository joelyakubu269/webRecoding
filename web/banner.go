package main

import (
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
}

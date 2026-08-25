package spellcheck

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Dictionary map[string]struct{}

func LoadDictionary(filename string) (Dictionary, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer file.Close()

	var dict = make(map[string]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dict[word] = struct{}{}
	}

	return dict, nil
}

package parser

import (
	"log"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	src, err := os.ReadFile("./testdata/unicode.txt")
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}

	tokens, err := Parse(src)

	log.Printf("%#v", tokens)
}

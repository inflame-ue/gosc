package parser

import (
	"fmt"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	src, err := os.ReadFile("./testdata/simple.txt")
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}

	tokens, err := Parse(src)

	fmt.Printf("%#v", tokens)
}

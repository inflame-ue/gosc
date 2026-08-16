package parser

import (
	"bytes"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "regenerate .golden files")

func TestParser(t *testing.T) {
	dir, err := os.ReadDir("testdata")
	if err != nil {
		t.Fatalf("reading directory: %v", err)
	}

	for _, entry := range dir {
		filename := filepath.Join("testdata", entry.Name())
		ext := filepath.Ext(entry.Name())

		if entry.IsDir() {
			continue
		}

		if ext == ".golden" {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			src, err := os.ReadFile(filename)
			if err != nil {
				t.Fatalf("reading src file: %v", err)
			}

			tokens, err := Parse(src)
			if err != nil {
				t.Fatal(err)
			}

			out, err := json.MarshalIndent(tokens, "", "  ")
			if err != nil {
				t.Fatalf("marhsalling tokens to JSON: %v", err)
			}

			goldenFile, _, _ := strings.Cut(filename, ext)
			goldenFile += ".golden"

			if *update {
				err = os.WriteFile(goldenFile, out, 0o644)
				if err != nil {
					t.Fatalf("writing golden file: %v", err)
				}
				return
			}

			golden, err := os.ReadFile(goldenFile)
			if err != nil {
				t.Fatalf("reading golden file: %v", err)
			}

			if !bytes.Equal(out, golden) {
				t.Error("parsed bytes do not equal to the golden bytes")
			}
		})
	}
}

func TestParserInvalid(t *testing.T) {
	dir, err := os.ReadDir("testdata/invalid")
	if err != nil {
		t.Fatalf("reading directory: %v", err)
	}

	for _, entry := range dir {
		filename := filepath.Join("testdata/invalid", entry.Name())

		if entry.IsDir() {
			continue
		}

		t.Run(entry.Name(), func(t *testing.T) {
			src, err := os.ReadFile(filename)
			if err != nil {
				t.Fatalf("reading file: %v", err)
			}

			_, err = Parse(src)
			if err == nil {
				t.Error("expected parsing error, got err = nil")
			}
		})
	}
}

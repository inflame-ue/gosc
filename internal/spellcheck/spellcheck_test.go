package spellcheck

import (
	"slices"
	"testing"

	"github.com/inflame-ue/gosc/internal/parser"
)

func TestCheck(t *testing.T) {
	dictionary := Dictionary{
		"hello": struct{}{},
	}

	tests := map[string]struct {
		tokens []parser.Token
		dict   Dictionary
		want   []Misspelling
	}{
		"no misspelled": {
			tokens: []parser.Token{
				parser.Token{Word: "hello", Line: 0, Col: 0},
			},
			dict: dictionary,
			want: nil,
		},
		"uppercase no misspelled": {
			tokens: []parser.Token{
				parser.Token{Word: "HELLO", Line: 0, Col: 0},
			},
			dict: dictionary,
			want: nil,
		},
		"misspelled": {
			tokens: []parser.Token{
				parser.Token{Word: "helo", Line: 0, Col: 0},
			},
			dict: dictionary,
			want: []Misspelling{
				Misspelling{Word: "helo", Line: 0, Col: 0},
			},
		},
		"uppercase misspelled": {
			tokens: []parser.Token{
				parser.Token{Word: "HELO", Line: 0, Col: 0},
			},
			dict: dictionary,
			want: []Misspelling{
				Misspelling{Word: "HELO", Line: 0, Col: 0},
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Check(tc.tokens, tc.dict)

			if !slices.Equal(tc.want, got) {
				t.Errorf("expected slice: %v, got: %v", tc.want, got)
			}
		})
	}
}

package spellcheck

import (
	"strings"

	"github.com/inflame-ue/gosc/internal/parser"
)

type Misspelling struct {
	Word      string
	Line, Col int
}

func Check(tokens []parser.Token, dict Dictionary) []Misspelling {
	var misspelled []Misspelling

	for _, token := range tokens {
		word := strings.ToLower(token.Word)
		if _, ok := dict[word]; ok {
			continue
		}

		misspelling := Misspelling{
			Word: token.Word,
			Line: token.Line,
			Col:  token.Col,
		}
		misspelled = append(misspelled, misspelling)
	}

	return misspelled
}

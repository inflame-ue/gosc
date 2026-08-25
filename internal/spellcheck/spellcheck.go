package spellcheck

import (
	"github.com/inflame-ue/gosc/internal/parser"
)

type Misspelling struct {
	Word string
	Line, Col      int
}

func Spellcheck(tokens []parser.Token, dict Dictionary) []Misspelling {
	var misspelled []Misspelling

	for _, token := range tokens {
		if _, ok := dict[token.Word]; ok {
			continue
		}

		misspelled = append(misspelled, Misspelling(token))
	}

	return misspelled
}

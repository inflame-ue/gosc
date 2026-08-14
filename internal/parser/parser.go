package parser

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

// Represents a single token in the source document
type Token struct {
	Word      string
	Line, Col int
}

// Parse accepts a slice of bytes that represent a source document.
// The characters in the source document must be of valid UTF-8 format.
//
//   - If src is a empty or zero-length slice, Parse returns nil for both values
//   - If src does not contain any letters, Parse also returns nil for both values
//   - Otherwise, Parse returns a slice of Tokens and a nil error
//
// In case of a parsing error Parse will return a nil slice and non-nill error.
func Parse(src []byte) ([]Token, error) {
	if len(src) == 0 {
		return nil, nil
	}

	if !utf8.Valid(src) {
		return nil, errors.New("err: source cannot contain invalid UTF-8 symbols")
	}

	var tokens []Token
	var buf []rune
	var line, col int

	for i, w := 0, 0; i < len(src); i += w {
		c, width := utf8.DecodeRune(src[i:])
		w = width

		if unicode.IsLetter(c) {
			buf = append(buf, c)
			continue
		}

		col++
		if c == '\n' {
			line++
			col = 0
		}

		if len(buf) == 0 {
			continue
		}

		token := Token{
			Word: string(buf),
			Line: line,
			Col:  col,
		}
		tokens = append(tokens, token)

		col += len(token.Word)
		buf = nil
	}

	token := Token{
		Word: string(buf),
		Line: line,
		Col:  col,
	}
	tokens = append(tokens, token)

	return tokens, nil
}

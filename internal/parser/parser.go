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
	var line int
	var col = 1

	for i, w := 0, 0; i < len(src); i += w {
		// the idiomatic way to handle non-uniform rune decodes
		c, width := utf8.DecodeRune(src[i:])
		w = width

		if unicode.IsLetter(c) {
			buf = append(buf, c)
			continue
		}

		// no letters accumulated on non-letter char hit; skip
		if len(buf) == 0 {
			continue
		}

		token := Token{
			Word: string(buf),
			Line: line,
			Col:  col,
		}
		tokens = append(tokens, token)

		// reset col and increment line if we hit /n
		if c == '\n' {
			line++
			col = 1
		}

		// type cast to rune to count chars, not bytes
		col += len([]rune(token.Word)) + 1
		buf = nil
	}

	// flush the last token from buf
	if len(buf) != 0 {
		token := Token{
			Word: string(buf),
			Line: line,
			Col:  col,
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}

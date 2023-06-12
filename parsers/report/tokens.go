// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

package report

import (
	"unicode"
)

type tokens []*token

func (t tokens) LineNo() int {
	if len(t) == 0 {
		return -1
	}
	return t[0].line
}

type token struct {
	line  int
	value string
}

// next returns the next token from the input.
// tokens are new-line, quoted text, or text.
func next(b []byte) (token, rest []byte) {
	for len(b) != 0 && b[0] != '\n' && unicode.IsSpace(rune(b[0])) {
		b = b[1:]
	}
	if len(b) == 0 {
		return nil, nil
	}

	token, rest = b[:1], b[1:]
	if token[0] == '\n' {
		return token, rest
	}

	if token[0] == '"' {
		for len(rest) != 0 && rest[0] != '\n' && rest[0] != token[0] {
			token, rest = append(token, rest[0]), rest[1:]
		}
		if len(rest) != 0 && rest[0] == token[0] {
			token, rest = append(token, rest[0]), rest[1:]
		} else {
			// ignore non-terminated quoted text for the moment
		}
		return token, rest
	}

	for len(rest) != 0 && !unicode.IsSpace(rune(rest[0])) {
		token, rest = append(token, rest[0]), rest[1:]
	}
	return token, rest
}

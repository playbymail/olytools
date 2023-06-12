// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

// Package report implements a parser for G4 format turn reports.
package report

import (
	"fmt"
	"io"
)

type Parser struct {
	input  []byte
	tokens tokens
}

// NewParser returns a new Parser parsing from b.
func NewParser(b []byte) (*Parser, error) {
	return &Parser{input: b, tokens: tokenizer(b)}, nil
}

func (p *Parser) Dump(w io.Writer) {
	printf := func(format string, args ...any) {
		_, _ = w.Write([]byte(fmt.Sprintf(format, args...)))
	}

	line := 1
	printf("%4d:\t", line)
	for _, t := range p.tokens {
		if t.value == "\n" {
			line++
			printf("\n%4d:\t", line)
			continue
		}
		printf(" «%s»", t.value)
	}
}

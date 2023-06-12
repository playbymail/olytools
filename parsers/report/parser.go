// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

package report

import (
	"bytes"
	"fmt"
)

func Parse(input []byte) (*Report, error) {
	_, err, rpt := tokenizer(input).expectReport()
	if err != nil {
		return rpt, err
	}
	return rpt, err
}

func tokenizer(input []byte) tokens {
	var t tokens
	line := 1
	for w, rest := next(input); w != nil; w, rest = next(rest) {
		// when we find a leading paren, see if we need to split it
		if w[0] == '(' && (len(w) > 1 && '0' <= w[1] && w[1] <= '9') {
			t = append(t, &token{line: line, value: "("})
			w = w[1:]
		}

		// when we find a leading bracket, closing bracket followed by a comma or period, we need to split it
		if bytes.HasPrefix(w, []byte{'['}) && (bytes.HasSuffix(w, []byte{']', ','}) || bytes.HasSuffix(w, []byte{']', '.'})) {
			t = append(t, &token{line: line, value: string(w[:len(w)-1])})
			w = w[len(w)-1:]
		}

		// when we find integers followed by punctuation, we must split them into two tokens.
		if '0' <= w[0] && w[0] <= '9' {
			nlen := 0
			for nlen < len(w) && '0' <= w[nlen] && w[nlen] <= '9' {
				nlen++
			}
			if nlen == len(w)-1 && (w[nlen] == '.' || w[nlen] == ',' || w[nlen] == '%') {
				// it is a number followed by punctuation
				t = append(t, &token{line: line, value: string(w[:nlen])})
				w = w[nlen:]
			}
		}

		t = append(t, &token{line: line, value: string(w)})
		if w[0] == '\n' {
			line++
		}
	}
	fmt.Printf("tokenizer: read %d tokens\n", len(t))

	return t
}

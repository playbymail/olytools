// olytools - tools to help play Olympia
// Copyright (C) 2023 Michael D Henderson. All rights reserved.

package report

import (
	"strconv"
)

func (t tokens) acceptEndOfInput() bool {
	return len(t) == 0
}

func (t tokens) acceptEndOfLine() bool {
	if len(t) == 0 {
		return false
	}
	if t[0].value == "\n" {
		return true
	}
	return false
}

func (t tokens) acceptInteger() bool {
	if len(t) == 0 {
		return false
	} else if t[0].value == "\n" {
		return false
	} else if _, err := strconv.Atoi(t[0].value); err != nil {
		return false
	}
	return true
}

func (t tokens) acceptLiteral(s string) bool {
	if len(t) == 0 {
		return false
	}
	return t[0].value == s
}

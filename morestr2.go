package morestr

import (
	"unicode/utf8"
)

func RuneLen(in string) int {
	return utf8.RuneCountInString(in)
}

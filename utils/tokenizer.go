package utils

import (
	"strings"
	"unicode"
)

func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = stopWordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

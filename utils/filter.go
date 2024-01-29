package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

func lowercaseFilter(tokens []string) []string {
	for i, token := range tokens {
		tokens[i] = strings.ToLower(token)
	}
	return tokens
}

func stopWordFilter(tokens []string) []string {
	var stopwords = map[string]struct{}{
		"a": {}, "an": {}, "and": {}, "are": {}, "as": {}, "at": {}, "be": {}, "but": {}, "by": {}, "i": {}, "if": {}, "in": {}, "into": {}, "is": {}, "it": {}, "no": {}, "not": {}, "of": {}, "on": {}, "or": {}, "such": {}, "that": {}, "the": {}, "their": {}, "then": {}, "there": {}, "these": {}, "they": {}, "this": {}, "to": {}, "was": {}, "will": {}, "with": {},
	}

	for i, token := range tokens {
		if _, ok := stopwords[token]; ok {
			tokens[i] = ""
		}
	}
	return tokens
}

func stemmerFilter(tokens []string) []string {
	for i, token := range tokens {
		tokens[i] = snowballeng.Stem(token, false)
	}
	return tokens
}

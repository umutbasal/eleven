package main

import "strings"

func lowercaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func stemEnglishFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = StemEnglish(token)
	}
	return r
}

var commonWords = map[string]bool{ // I wish Go had built-in sets.
	"the":  true,
	"be":   true,
	"to":   true,
	"of":   true,
	"and":  true,
	"a":    true,
	"in":   true,
	"that": true,
	"have": true,
	"i":    true,
	"it":   true,
}

func commonWordsFilter(tokens []string) []string {
	r := make([]string, 0)
	for _, token := range tokens {
		if !commonWords[token] {
			r = append(r, token)
		}
	}
	return r
}

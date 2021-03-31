package main

import "strings"

func lowercaseFilter(tokens []string) []string {
	r := make([]string, cap(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func stemEnglishFilter(tokens []string) []string {
	r := make([]string, cap(tokens))
	for i, token := range tokens {
		r[i] = StemEnglish(token)
	}
	return r
}

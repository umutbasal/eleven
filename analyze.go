package main

func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowercaseFilter(tokens)
	tokens = commonWordsFilter(tokens)
	tokens = stemEnglishFilter(tokens)
	return tokens
}

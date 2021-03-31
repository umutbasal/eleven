package main

import "strings"

func search(docs []document, term string) []document {
	var r []document
	for _, d := range docs {
		if strings.Contains(d.Title, term) {
			r = append(r, d)
		}
	}
	return r
}

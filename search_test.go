package main

import "testing"

func TestSearch(t *testing.T) {
	doc, _ := loadDocuments("example.xml")

	actual := search(doc, "Glock")[0].URL
	expected := "https://en.wikipedia.org/wiki/Kit-Cat_Glock"

	if actual != expected {
		t.Errorf("actual = %v, expected = %v", actual, expected)
	}
}

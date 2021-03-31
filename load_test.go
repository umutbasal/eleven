package main

import (
	"log"
	"strings"
	"testing"
)

func TestLoadDocuments(t *testing.T) {
	doc, err := loadDocuments("example.xml")

	if err != nil {
		t.Fatalf("Failed to load")
	}

	actual := doc[0].Title
	expected := "Klock"
	fail := "Whoop"

	log.Printf("%v", doc)
	if !strings.Contains(actual, expected) {
		t.Errorf("expected %s in title , actual %s", expected, actual)
	}

	if strings.Contains(actual, fail) {
		t.Errorf("did not expected %s to be in %s", fail, actual)
	}

	if cap(doc) != 2 {
		t.Errorf("expected cap 2, but got %d", cap(doc))
	}
}

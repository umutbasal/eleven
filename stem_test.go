package main

import (
	"fmt"
	"os"
	"testing"
)

func TestNewStemmer(t *testing.T) {
	tests := []struct {
		algorithm string
		word      string
		stem      string
	}{
		{"english", "accumulations", "accumul"},
		{"english", "eating", "eat"},
		{"turkish", "kalelerimizdekilerden", "kale"},
		{"turkish", "çocuğuymuşumcasına", "çocuk"},
	}
	for _, tc := range tests {
		s, err := NewStemmer(tc.algorithm)
		if err != nil {
			fmt.Println("Error creating stemmer: " + err.Error())
			os.Exit(1)
		}
		actual := s.Stem(tc.word)
		expected := tc.stem

		if actual != expected {
			t.Errorf("expected %s, got %s", expected, actual)
		}
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	actual := tokenize("hello world how are you?")
	expected := []string{"hello", "world", "how", "are", "you"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual %v, Expected %v", actual, expected)
	}
}

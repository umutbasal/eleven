package main

import (
	"reflect"
	"testing"
)

func Test_analyze(t *testing.T) {
	tests := []struct {
		text     string
		expected []string
	}{
		{"Friends with benefits. ", []string{"friend", "with", "benefit"}},
		{"a tomato went to bar and spoke with girls", []string{"tomato", "went", "bar", "spoke", "with", "girl"}},
	}
	for _, tc := range tests {
		if actual := analyze(tc.text); !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("analyze() = %v, want %v", actual, tc.expected)
		}
	}
}

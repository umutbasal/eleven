package main

import (
	"reflect"
	"testing"
)

func TestLowerCaseFilter(t *testing.T) {
	actual := lowercaseFilter([]string{"Hello", "World"})
	expected := []string{"hello", "world"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual %#v, expected %#v", actual, expected)
	}

}

func TestStemEnglishFilter(t *testing.T) {
	actual := stemEnglishFilter([]string{"eating", "dogs"})
	expected := []string{"eat", "dog"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual %#v, expected %#v", actual, expected)
	}

}

func TestCommonWordsFilter(t *testing.T) {
	actual := commonWordsFilter([]string{"i", "have", "that", "eating", "problem"})
	expected := []string{"eating", "problem"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Actual %#v, expected %#v", actual, expected)
	}

}

package main

import (
	"reflect"
	"testing"
)

func Test_index_feed(t *testing.T) {
	actual := make(index)
	actual.feed([]document{{ID: 1, Text: "A donut on a glass plate. Only the donuts."}})
	actual.feed([]document{{ID: 2, Text: "donut is a donut"}})

	expected := index{
		"donut": {1, 2},
		"glass": {1},
		"is":    {2},
		"on":    {1},
		"only":  {1},
		"plate": {1},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %#v, actual %#v", expected, actual)
	}
}

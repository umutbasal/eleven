package main

import (
	"reflect"
	"testing"
)

func TestEleven(t *testing.T) {
	err := indexAndSave("example.xml")
	if err != nil {
		t.Errorf("problem with indexing and saving %v", err)
	}

	table, err := loadIndexedTable("example.xml.json")
	if err != nil {
		t.Errorf("problem with loadIndexedTable %v", err)
	}
	actual := table.search("glock")
	expected := [][]int{{1}}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %v, expected %v", actual, expected)
	}
}

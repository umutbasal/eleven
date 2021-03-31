package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

func save(table index) {
	file, _ := json.MarshalIndent(table, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}

func indexAndSave(filename string) (err error) {
	doc, err := loadDocuments(filename)
	if err != nil {
		return errors.New("failed to load")
	}
	table := make(index)
	table.feed(doc)
	save(table)
	return nil
}

func loadIndexedTable(filename string) (indexTable index, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &indexTable)
	return indexTable, err
}

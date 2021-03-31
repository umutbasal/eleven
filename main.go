package main

import (
	"fmt"
	"log"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func main() {
	defer elapsed("load")()

	indexAndSave("example.xml")

	table, err := loadIndexedTable("test.json")

	if err != nil {
		log.Fatalf("couldn't load' %v", err)
	}
	log.Printf("results: %v", table.search("glock"))
}

package main

import (
	"fmt"
	"log"

	"myproject/datafile"
)

func main() {
	texts, err := datafile.GetStrings(votes.txt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(texts)
}

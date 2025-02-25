package main

import (
	"fmt"
	"log"

	"myproject/datafile"
)

func main() {
	texts, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}
	var names []string
	var counts []int
	for _, text := range texts {
		matched := false
		for i, name := range names {
			if name == text {
				counts[i]++
				matched = true
			}

		}
		if matched == false {
			names = append(names, text)
			counts = append(counts, 1)
		}
	}
	for i, name := range names {
		fmt.Printf("%s: %d\n", name, counts[i])
	}
}

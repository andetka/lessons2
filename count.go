package main

import (
	"fmt"
	"log"
	"sort"

	"myproject/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	var alpha []string
	if err != nil {
		log.Fatal(err)
	}
	counts := make(map[string]int)
	for _, line := range lines {
		counts[line]++
	}
	for name := range counts {
		alpha = append(alpha, name)
	}
	sort.Strings(alpha)
	for _, word := range alpha {
		fmt.Println(word, counts[word])
	}
}

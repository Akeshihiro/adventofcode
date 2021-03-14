package main

import (
	"adventofcode2020/internal/pkg/day6"
	"log"
)

func main() {
	lines, err := day6.ReadInputFile("data/day6/input.txt")
	if err != nil {
		log.Panic(err)
	}

	groups := day6.ExtractGroups(lines)
	sum := 0
	for _, group := range groups {
		answers := day6.UniqueAnswersOfGroup(group)
		sum += len(answers)
	}
	log.Printf("sum of answers: %v", sum)
}

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
	sumByAll := 0
	for _, groupAnswers := range groups {
		answers := day6.UniqueAnswersOfGroup(groupAnswers)
		sum += len(answers)

		answersByAll := day6.UniqueAnswersAnsweredByAllGroupMembers(groupAnswers)
		sumByAll += len(answersByAll)
	}
	log.Printf("sum of answers: %v", sum)
	log.Printf("sum of answers answered by all group members: %v", sumByAll)
}

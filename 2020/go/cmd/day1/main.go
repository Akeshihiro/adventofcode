package main

import (
	"adventofcode2020/internal/pkg/day1"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := day1.ReadInput(file)
	if err != nil {
		log.Fatal(err)
	}

	s1, s2, err := day1.FindSummandsOf(2020, data)
	if err != nil {
		log.Fatal(err)
	}

	result := s1 * s2
	log.Println(result)
}

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

	s1, s2, err := day1.FindTwoSummandsOf(2020, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("s1 * s2 = %v", s1*s2)

	s1, s2, s3, err := day1.FindThreeSummandsOf(2020, data)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("s1 * s2 * s3 = %v", s1*s2*s3)
}

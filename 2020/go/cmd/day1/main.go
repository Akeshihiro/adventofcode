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
	defer file.Close()

	data, err := day1.ReadInput(file)
	if err != nil {
		log.Fatal(err)
	}

	result, err := day1.FindSummandsOf(2020, data, 2)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("s1 * s2 = %v", result[0]*result[1])

	result, err = day1.FindSummandsOf(2020, data, 3)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("s1 * s2 * s3 = %v", result[0]*result[1]*result[2])
}

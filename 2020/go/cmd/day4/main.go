package main

import (
	"adventofcode2020/internal/pkg/day4"
	"log"
)

func main() {
	data, err := day4.ReadInputFile("data/day4/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	count := day4.CountValidPassports(data)

	log.Printf("valid passports: %v\n", count)
}

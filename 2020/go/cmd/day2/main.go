package main

import (
	"adventofcode2020/internal/pkg/day2"
	"log"
)

func main() {
	data, err := day2.ReadInputFile("data/day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// valid, invalid := day2.ValidatePasswords(data)
	valid, invalid := day2.ValidatePasswordsByTobogganPolicies(data)
	log.Printf("valid = %v, invalid = %v", valid, invalid)
}

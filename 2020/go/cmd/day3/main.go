package main

import (
	"adventofcode2020/internal/pkg/day3"
	"log"
)

func main() {
	matrix, err := day3.ReadInputFile("data/day3/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	foundTrees := day3.CountTreesByMoving(matrix, 3, 1)

	log.Println(foundTrees)
}

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

	foundTrees, err := day3.CountTreesByMoving(matrix,
		1, 1,
		3, 1,
		5, 1,
		7, 1,
		1, 2,
	)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(foundTrees)

	product := 1
	for _, v := range foundTrees {
		product *= v
	}

	log.Println(product)
}

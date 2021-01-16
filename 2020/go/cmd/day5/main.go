package main

import (
	"adventofcode2020/internal/pkg/day5"
	"log"
)

func main() {
	seatnums, err := day5.ReadInputFile("data/day5/input.txt")
	if err != nil {
		log.Panic(err)
	}

	highestSeatID := 0
	for _, v := range seatnums {
		row, err := day5.DecodeRow(v)
		if err != nil {
			log.Print(err)
			continue
		}

		col, err := day5.DecodeCol(v)
		if err != nil {
			log.Print(err)
			continue
		}

		seatID := row*8 + col
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}

	log.Println(highestSeatID)
}

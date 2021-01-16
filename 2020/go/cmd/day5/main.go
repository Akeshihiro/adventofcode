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

	seatIDs := []int{}
	highestSeatID := -1
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

		seatIDs = append(seatIDs, seatID)
	}

	log.Println(highestSeatID)

	// Part 2
	missingSeats := day5.FindMissingSeats(seatIDs)
	log.Print(missingSeats)
}

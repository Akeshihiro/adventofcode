package day5

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

const (
	firstRow = 0
	lastRow  = 127

	firstCol = 0
	lastCol  = 7
)

var (
	seatcodePattern = regexp.MustCompile(`^([BF]{7})([LR]{3})$`)
)

func ReadInputFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ReadInput(file)
}

func ReadInput(r io.Reader) ([]string, error) {
	result := []string{}

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	content := string(b)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

		if !seatcodePattern.MatchString(line) {
			log.Printf("seat number '%v' is not valid, skipping", line)
			continue
		}

		result = append(result, line)
	}

	return result, nil
}

func DecodeRow(seatcode string) (int, error) {
	if !seatcodePattern.MatchString(seatcode) {
		return -1, fmt.Errorf("invalid seat number '%v", seatcode)
	}

	rowcode := seatcodePattern.FindStringSubmatch(seatcode)[1]
	start := firstRow
	end := lastRow
	for _, v := range rowcode {
		v := unicode.ToUpper(v)
		switch v {
		case 'F':
			end = (start + end) / 2
		case 'B':
			start = int(math.Ceil(float64(start+end) / 2))
		}
	}

	return start, nil
}

func DecodeCol(seatcode string) (int, error) {
	if !seatcodePattern.MatchString(seatcode) {
		return -1, fmt.Errorf("invalid seat number '%v", seatcode)
	}

	colcode := seatcodePattern.FindStringSubmatch(seatcode)[2]
	start := firstCol
	end := lastCol
	for _, v := range colcode {
		v := unicode.ToUpper(v)
		switch v {
		case 'L':
			end = (start + end) / 2
		case 'R':
			start = int(math.Ceil(float64(start+end) / 2))
		}
	}

	return start, nil
}

func FindMissingSeats(seats []int) []int {
	if len(seats) == 0 {
		return nil
	}

	// We need a sorted slice but the original slice must stay untouched!
	ids := make([]int, len(seats))
	copy(ids, seats)
	sort.Ints(ids)

	result := []int{}
	lastCheckedSeatID := ids[0]
	for i := 1; i < len(ids); i++ {
		id := ids[i]

		if lastCheckedSeatID+1 != id {
			for j := lastCheckedSeatID + 1; j < id; j++ {
				result = append(result, j)
			}
		}

		lastCheckedSeatID = id
	}

	return result
}

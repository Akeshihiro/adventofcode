package day3

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

func ReadInputFile(filename string) ([][]rune, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadInput(f)
}

func ReadInput(r io.Reader) ([][]rune, error) {
	reader := bufio.NewReader(r)

	result := [][]rune{}

	for {
		isLastLine := false
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				isLastLine = true
			} else {
				return nil, err
			}
		}

		line = strings.TrimSpace(line)
		result = append(result, []rune(line))

		if isLastLine {
			break
		}
	}

	return result, nil
}

func CountTreesByMoving(matrix [][]rune, moves ...int) ([]int, error) {
	l := len(moves)
	if l%2 != 0 {
		return nil, errors.New("too few or too much movement arguments; must be an even length of movement arguments as they are x:y pairs")
	}

	pairs := l / 2
	counters := []int{}

	rows := len(matrix)
	cols := 0
	if rows > 0 {
		cols = len(matrix[0])
	}

	for i := 0; i < pairs; i++ {
		x := moves[i*2]
		y := moves[i*2+1]
		counters = append(counters, 0)

		for r, c := 0, 0; r < rows; r, c = r+y, (c+x)%cols {
			if matrix[r][c] == '#' {
				counters[i]++
			}
		}
	}

	return counters, nil
}

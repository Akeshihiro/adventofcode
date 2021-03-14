package day1

import (
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	ErrNilInputData  = errors.New("input data == nil")
	ErrNoResultFound = errors.New("no result found")
)

func ReadInput(r io.Reader) ([]int, error) {
	if r == nil {
		return nil, ErrNilInputData
	}

	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	data := string(bytes)
	parts := strings.Split(data, "\n")

	result := []int{}
	for _, v := range parts {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}

		result = append(result, i)
	}

	return result, nil
}

func FindSummandsOf(sum int, data []int, amount uint) ([]int, error) {
	if data == nil {
		return nil, ErrNilInputData
	}

	if len(data) == 0 {
		return nil, ErrNoResultFound
	}

	return findSummandsOf(sum, data, amount, 0, nil)
}

func findSummandsOf(sum int, data []int, level uint, accSum int, accData []int) ([]int, error) {
	if level <= 0 {
		if accSum == sum {
			return accData, nil
		}

		return nil, ErrNoResultFound
	}

	for _, v := range data {
		result, err := findSummandsOf(sum, data, level-1, accSum+v, append(accData, v))
		if err != nil {
			continue
		}

		return result, nil
	}

	return nil, ErrNoResultFound
}

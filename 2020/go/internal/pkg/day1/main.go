package day1

import (
	"errors"
	"io"
	"io/ioutil"
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

	bytes, err := ioutil.ReadAll(r)
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

func FindTwoSummandsOf(sum int, data []int) (int, int, error) {
	if data == nil {
		return 0, 0, ErrNilInputData
	}

	if len(data) == 0 {
		return 0, 0, ErrNoResultFound
	}

	for i, iv := range data {
		for _, jv := range data[i:] {
			if iv+jv == sum {
				return iv, jv, nil
			}
		}
	}

	return 0, 0, ErrNoResultFound
}

func FindThreeSummandsOf(sum int, data []int) (int, int, int, error) {
	if data == nil {
		return 0, 0, 0, ErrNilInputData
	}

	if len(data) == 0 {
		return 0, 0, 0, ErrNoResultFound
	}

	for _, iv := range data {
		for _, jv := range data {
			for _, kv := range data {
				if iv+jv+kv == sum {
					return iv, jv, kv, nil
				}
			}
		}
	}

	return 0, 0, 0, ErrNoResultFound
}

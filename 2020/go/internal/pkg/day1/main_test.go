package day1_test

import (
	"adventofcode2020/internal/pkg/day1"
	"strings"
	"testing"
)

func TestReadInput_WhenInputReaderIsNil_ThenReturnAnError(t *testing.T) {
	_, err := day1.ReadInput(nil)
	if err == nil {
		t.Errorf("err == nil")
	}
}

func TestReadInput_WhenInputIsEmpty_ThenReturnAnEmptyResult(t *testing.T) {
	reader := strings.NewReader("")
	data, err := day1.ReadInput(reader)
	if err != nil {
		t.Errorf("err != nil, %v", err)
	}

	if data == nil {
		t.Errorf("data == nil")
	}

	if len(data) > 0 {
		t.Errorf("len(data) > 0, %v", len(data))
	}
}

func TestReadInput_WhenInputIsNotEmpty_ThenReturnASliceWithData(t *testing.T) {
	reader := strings.NewReader(`1
2
3`)
	data, err := day1.ReadInput(reader)
	if err != nil {
		t.Errorf("err != nil, %v", err)
	}

	if data == nil {
		t.Errorf("data == nil")
	}

	l := len(data)
	if l != 3 {
		t.Errorf("len(data) != 3, %v", l)
	}
}

func TestReadInput_WhenInputHasEmptyLines_ThenIgnoreEmptyLines(t *testing.T) {
	reader := strings.NewReader(`1
2

3
`)
	data, err := day1.ReadInput(reader)
	if err != nil {
		t.Errorf("err != nil, %v", err)
	}

	if data == nil {
		t.Errorf("data == nil")
	}

	l := len(data)
	if l != 3 {
		t.Errorf("len(data) != 3, %v", l)
	}
}

func TestFindSummandOf_WhenDataIsNil_ThenThrowAnError(t *testing.T) {
	_, _, err := day1.FindSummandsOf(0, nil)
	if err != day1.ErrNilInputData {
		t.Errorf("err != ErrNilInputData")
	}
}

func TestFindSummandOf_WhenDataIsEmpty_ThenThrowAnError(t *testing.T) {
	_, _, err := day1.FindSummandsOf(0, []int{})
	if err != day1.ErrNoResultFound {
		t.Errorf("err != ErrNoResultFound")
	}
}

func TestFindSummandOf_Return1721And299(t *testing.T) {
	data := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	s1, s2, err := day1.FindSummandsOf(2020, data)
	if err != nil {
		t.Errorf("err != nil")
	}
	if s1 != 1721 && s2 != 299 {
		t.Errorf("s1 != 1721 or s2 != 299, s1=%v, s2=%v", s1, s2)
	}
}

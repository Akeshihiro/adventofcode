package day1_test

import (
	"adventofcode2020/internal/pkg/day1"
	"reflect"
	"strings"
	"testing"
)

func TestReadInput_WhenInputReaderIsNil_ThenReturnErrNilInputData(t *testing.T) {
	_, err := day1.ReadInput(nil)
	if err != day1.ErrNilInputData {
		t.Fatalf("expected %v, got %v", day1.ErrNilInputData, err)
	}
}

func TestReadInput_WhenInputIsEmpty_ThenReturnAnEmptyResult(t *testing.T) {
	reader := strings.NewReader("")
	result, _ := day1.ReadInput(reader)
	expected := []int{}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestReadInput_WhenInputIsNotEmpty_ThenReturnASliceWithData(t *testing.T) {
	reader := strings.NewReader(`1
2
3`)
	result, _ := day1.ReadInput(reader)
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestReadInput_WhenInputHasEmptyLines_ThenIgnoreEmptyLines(t *testing.T) {
	reader := strings.NewReader(`1
2

3
`)
	result, _ := day1.ReadInput(reader)
	expected := []int{1, 2, 3}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestFindSummandOf_WhenDataIsNil_ThenReturnErrNilInputData(t *testing.T) {
	_, err := day1.FindSummandsOf(0, nil, 0)
	if err != day1.ErrNilInputData {
		t.Fatalf("expected %v, got %v", day1.ErrNilInputData, err)
	}
}

func TestFindSummandOf_WhenDataIsEmpty_ThenReturnErrNoResultFound(t *testing.T) {
	_, err := day1.FindSummandsOf(0, []int{}, 0)
	if err != day1.ErrNoResultFound {
		t.Fatalf("expected %v, got %v", day1.ErrNoResultFound, err)
	}
}

func TestFindSummandOf_WhenAmountOfSummandsShouldBe2_ThenReturn1721And299(t *testing.T) {
	data := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	result, _ := day1.FindSummandsOf(2020, data, 2)
	expected := []int{1721, 299}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestFindSummandOf_WhenAmountOfSummandsShouldBe3_ThenReturn979And366And675(t *testing.T) {
	data := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	result, _ := day1.FindSummandsOf(2020, data, 3)
	expected := []int{979, 366, 675}

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

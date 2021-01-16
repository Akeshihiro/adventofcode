package day5_test

import (
	"adventofcode2020/internal/pkg/day5"
	"reflect"
	"testing"
)

func TestDecodeRow_WhenSeatcodeIsFBFBBFFRLR_ThenReturn44(t *testing.T) {
	input := "FBFBBFFRLR"
	expected := 44
	result, _ := day5.DecodeRow(input)

	if result != expected {
		t.Fatalf("expected '%v', got '%v'", expected, result)
	}
}

func TestDecodeCol_WhenSeatcodeIsFBFBBFFRLR_ThenReturn5(t *testing.T) {
	input := "FBFBBFFRLR"
	expected := 5
	result, _ := day5.DecodeCol(input)

	if result != expected {
		t.Fatalf("expected '%v', got '%v'", expected, result)
	}
}

func TestFindMissingSeats_WhenThereIsAGap_ThenReturnTheMissingIDs(t *testing.T) {
	input := []int{1, 5}
	expected := []int{2, 3, 4}
	result := day5.FindMissingSeats(input)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected '%v', got '%v'", expected, result)
	}
}

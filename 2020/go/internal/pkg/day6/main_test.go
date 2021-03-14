package day6_test

import (
	"adventofcode2020/internal/pkg/day6"
	"reflect"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := strings.NewReader(`abc

a
b
c

ab
ac

a
a
a
a

b`)
	expected := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}
	result, _ := day6.ReadInput(input)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected '%v', got '%v'", expected, result)
	}
}

func TestExtractGroups(t *testing.T) {
	input := []string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}
	expected := [][]string{
		{"abc"},
		{"a", "b", "c"},
		{"ab", "ac"},
		{"a", "a", "a", "a"},
		{"b"},
	}
	result := day6.ExtractGroups(input)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected '%v', got '%v'", expected, result)
	}
}

func TestUniqueAnswersOfGroup(t *testing.T) {
	input := []string{
		"a",
		"az",
		"c",
		"cz",
	}
	expected := "acz"
	result := day6.UniqueAnswersOfGroup(input)

	if result != expected {
		t.Fatalf("expected '%v', got '%v'", expected, result)
	}
}

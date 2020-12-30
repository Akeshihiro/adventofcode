package day2_test

import (
	"adventofcode2020/internal/pkg/day2"
	"reflect"
	"strings"
	"testing"
)

func TestReadInput_WhenInputHasData_ThenReturnPasswordData(t *testing.T) {
	input := `1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc`
	reader := strings.NewReader(input)
	result, _ := day2.ReadInput(reader)
	expected := []day2.PasswordWithPolicy{
		{Password: "abcde", Sequence: "a", Min: 1, Max: 3},
		{Password: "cdefg", Sequence: "b", Min: 1, Max: 3},
		{Password: "ccccccccc", Sequence: "c", Min: 2, Max: 9},
	}

	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestValidatePasswords_WhenInputHas2ValidAnd1InvalidPasswords_ThenReturnValidEquals2AndInvalidEquals1(t *testing.T) {
	input := []day2.PasswordWithPolicy{
		{Password: "abcde", Sequence: "a", Min: 1, Max: 3},
		{Password: "cdefg", Sequence: "b", Min: 1, Max: 3},
		{Password: "ccccccccc", Sequence: "c", Min: 2, Max: 9},
	}
	valid, invalid := day2.ValidatePasswords(input)
	expectedValid, expectedInvalid := 2, 1

	if expectedValid != valid || expectedInvalid != invalid {
		t.Fatalf("expected valid=%v invalid=%v, got valid=%v invalid=%v", expectedValid, expectedInvalid, valid, invalid)
	}
}

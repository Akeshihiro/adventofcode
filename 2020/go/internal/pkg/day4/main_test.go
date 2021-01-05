package day4_test

import (
	"adventofcode2020/internal/pkg/day4"
	"reflect"
	"strings"
	"testing"
)

func TestReadInput(t *testing.T) {
	input := strings.NewReader(`ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`)
	expected := []day4.Passport{
		{
			BirthYearh:     "1937",
			IssueYear:      "2017",
			ExpirationYear: "2020",
			Height:         "183cm",
			HairColor:      "#fffffd",
			EyeColor:       "gry",
			PassportID:     "860033327",
			CountryID:      "147",
		},
		{
			BirthYearh:     "1929",
			IssueYear:      "2013",
			ExpirationYear: "2023",
			Height:         "",
			HairColor:      "#cfa07d",
			EyeColor:       "amb",
			PassportID:     "028048884",
			CountryID:      "350",
		},
		{
			BirthYearh:     "1931",
			IssueYear:      "2013",
			ExpirationYear: "2024",
			Height:         "179cm",
			HairColor:      "#ae17e1",
			EyeColor:       "brn",
			PassportID:     "760753108",
			CountryID:      "",
		},
		{
			BirthYearh:     "",
			IssueYear:      "2011",
			ExpirationYear: "2025",
			Height:         "59in",
			HairColor:      "#cfa07d",
			EyeColor:       "brn",
			PassportID:     "166559648",
			CountryID:      "",
		},
	}
	result, _ := day4.ReadInput(input)

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestCountValidPassports(t *testing.T) {
	input := []day4.Passport{
		{
			BirthYearh:     "1937",
			IssueYear:      "2017",
			ExpirationYear: "2020",
			Height:         "183cm",
			HairColor:      "#fffffd",
			EyeColor:       "gry",
			PassportID:     "860033327",
			CountryID:      "147",
		},
		{
			BirthYearh:     "1929",
			IssueYear:      "2013",
			ExpirationYear: "2023",
			Height:         "",
			HairColor:      "#cfa07d",
			EyeColor:       "amb",
			PassportID:     "028048884",
			CountryID:      "350",
		},
		{
			BirthYearh:     "1931",
			IssueYear:      "2013",
			ExpirationYear: "2024",
			Height:         "179cm",
			HairColor:      "#ae17e1",
			EyeColor:       "brn",
			PassportID:     "760753108",
			CountryID:      "",
		},
		{
			BirthYearh:     "",
			IssueYear:      "2011",
			ExpirationYear: "2025",
			Height:         "59in",
			HairColor:      "#cfa07d",
			EyeColor:       "brn",
			PassportID:     "166559648",
			CountryID:      "",
		},
	}
	expected := 2
	result := day4.CountValidPassports(input)

	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

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

func TestCountValidPassportsBySimpleValidation(t *testing.T) {
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
	result := day4.CountValidPassportsBySimpleValidation(input)

	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestCountValidPassportsByStrictValidation_WhenAllPassportsAreInvalid_ThenReturn0(t *testing.T) {
	input := strings.NewReader(`eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007`)
	data, _ := day4.ReadInput(input)
	expected := 0
	result := day4.CountValidPassportsByStrictValidation(data)

	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

func TestCountValidPassportsByStrictValidation_WhenThereAre4ValidPassports_ThenReturn4(t *testing.T) {
	input := strings.NewReader(`pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`)
	data, _ := day4.ReadInput(input)
	expected := 4
	result := day4.CountValidPassportsByStrictValidation(data)

	if result != expected {
		t.Fatalf("expected %v, got %v", expected, result)
	}
}

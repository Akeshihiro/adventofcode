package day4

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	BirthYearh     string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func ReadInputFile(filename string) ([]Passport, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return ReadInput(f)
}

func ReadInput(r io.Reader) ([]Passport, error) {
	reader := bufio.NewReader(r)
	result := []Passport{}

	var pp *Passport

	isEOF := false
	for !isEOF {
		line, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				isEOF = true
			} else {
				return nil, err
			}
		}

		line = strings.TrimSpace(line)

		if len(line) == 0 {
			// Empty line => persist the last parsed passport, if exists, ignore line otherwise
			if pp != nil {
				result = append(result, *pp)
				pp = nil
			}
		}

		if pp == nil {
			pp = &Passport{}
		}

		fields := strings.Fields(line)
		for _, f := range fields {
			fields := strings.Split(f, ":")
			if len(fields) != 2 {
				// Ignore wrong key-value-pairs for now
			}

			switch strings.ToLower(fields[0]) {
			case "byr":
				pp.BirthYearh = fields[1]
			case "iyr":
				pp.IssueYear = fields[1]
			case "eyr":
				pp.ExpirationYear = fields[1]
			case "hgt":
				pp.Height = fields[1]
			case "hcl":
				pp.HairColor = fields[1]
			case "ecl":
				pp.EyeColor = fields[1]
			case "pid":
				pp.PassportID = fields[1]
			case "cid":
				pp.CountryID = fields[1]
			default:
				// Ignore unknow fields for now
			}
		}

		if isEOF {
			// End of file => persist the last parsed passport, if exists
			if pp != nil {
				result = append(result, *pp)
				pp = nil
			}
		}
	}

	return result, nil
}

func CountValidPassportsBySimpleValidation(passports []Passport) int {
	counter := 0

	for _, p := range passports {
		if isPassportValidBySimpleValidation(p) {
			counter++
		}
	}

	return counter
}

func isPassportValidBySimpleValidation(p Passport) bool {
	return !(len(p.BirthYearh) == 0 ||
		len(p.IssueYear) == 0 ||
		len(p.ExpirationYear) == 0 ||
		len(p.HairColor) == 0 ||
		len(p.EyeColor) == 0 ||
		len(p.Height) == 0 ||
		len(p.PassportID) == 0)
}

func CountValidPassportsByStrictValidation(passports []Passport) int {
	counter := 0

	for _, p := range passports {
		if isPassportValidByStrictValidation(p) {
			counter++
		}
	}

	return counter
}

func isPassportValidByStrictValidation(p Passport) bool {
	return isBirthYearValid(p.BirthYearh) &&
		isIssueYearValid(p.IssueYear) &&
		isExpirationYearValid(p.ExpirationYear) &&
		isHairColorValid(p.HairColor) &&
		isEyeColorValid(p.EyeColor) &&
		isHeightValid(p.Height) &&
		isPassportIDValid(p.PassportID)
}

func isBirthYearValid(s string) bool {
	year, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return year >= 1920 && year <= 2002
}

func isIssueYearValid(s string) bool {
	year, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return year >= 2010 && year <= 2020
}

func isExpirationYearValid(s string) bool {
	year, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return year >= 2020 && year <= 2030
}

func isHairColorValid(s string) bool {
	s = strings.ToLower(s)
	pattern := regexp.MustCompile("^#[a-f0-9]{6}$")

	return pattern.MatchString(s)
}

func isEyeColorValid(s string) bool {
	s = strings.ToLower(s)
	pattern := regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")

	return pattern.MatchString(s)
}

func isHeightValid(s string) bool {
	s = strings.ToLower(s)
	pattern := regexp.MustCompile("^([0-9]{2,3})(cm|in)$")

	if !pattern.MatchString(s) {
		return false
	}

	matches := pattern.FindStringSubmatch(s)
	if len(matches) != 3 {
		return false
	}

	height, err := strconv.Atoi(matches[1])
	if err != nil {
		return false
	}

	switch matches[2] {
	case "cm":
		if height < 150 || height > 193 {
			return false
		}
	case "in":
		if height < 59 || height > 76 {
			return false
		}
	}

	return true
}

func isPassportIDValid(s string) bool {
	pattern := regexp.MustCompile("^[0-9]{9}$")

	return pattern.MatchString(strings.ToLower(s))
}

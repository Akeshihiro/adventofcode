package day4

import (
	"bufio"
	"errors"
	"io"
	"os"
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

func CountValidPassports(passports []Passport) int {
	counter := 0

	for _, p := range passports {
		if isPassportValid(p) {
			counter++
		}
	}

	return counter
}

func isPassportValid(p Passport) bool {
	return !(len(p.BirthYearh) == 0 ||
		len(p.IssueYear) == 0 ||
		len(p.ExpirationYear) == 0 ||
		len(p.HairColor) == 0 ||
		len(p.EyeColor) == 0 ||
		len(p.Height) == 0 ||
		len(p.PassportID) == 0)
}

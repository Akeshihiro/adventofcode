package day2

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type PasswordWithPolicy struct {
	Password string
	Sequence string
	Min      uint
	Max      uint
}

func ReadInputFile(file string) ([]PasswordWithPolicy, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	return ReadInput(f)
}

func ReadInput(r io.Reader) ([]PasswordWithPolicy, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	s := strings.TrimSpace(string(b))

	return parsePasswords(s)
}

func parsePasswords(s string) ([]PasswordWithPolicy, error) {
	lines := strings.Split(s, "\n")

	result := []PasswordWithPolicy{}

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		p, err := parsePassword(l)
		if err != nil {
			log.Println(err)
			continue
		}

		result = append(result, p)
	}

	return result, nil
}

func parsePassword(s string) (PasswordWithPolicy, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return PasswordWithPolicy{}, fmt.Errorf("empty line")
	}

	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return PasswordWithPolicy{}, fmt.Errorf("invalid format '%v'", s)
	}

	password := strings.TrimSpace(parts[1])

	parts = strings.Split(strings.TrimSpace(parts[0]), " ")
	if len(parts) != 2 {
		return PasswordWithPolicy{}, fmt.Errorf("invalid format '%v'", s)
	}

	sequence := strings.TrimSpace(parts[1])

	parts = strings.Split(strings.TrimSpace(parts[0]), "-")
	if len(parts) != 2 {
		return PasswordWithPolicy{}, fmt.Errorf("invalid format '%v'", s)
	}

	min, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return PasswordWithPolicy{}, fmt.Errorf("invalid format '%v'", s)
	}

	max, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return PasswordWithPolicy{}, fmt.Errorf("invalid format '%v'", s)
	}

	return PasswordWithPolicy{
		Password: password,
		Sequence: sequence,
		Min:      uint(min),
		Max:      uint(max),
	}, nil
}

func ValidatePasswords(passwords []PasswordWithPolicy) (int, int) {
	valid := 0
	invalid := 0

	for _, p := range passwords {
		c := strings.Count(p.Password, p.Sequence)
		if c < int(p.Min) || c > int(p.Max) {
			invalid++
		} else {
			valid++
		}
	}

	return valid, invalid
}

func ValidatePasswordsByTobogganPolicies(passwords []PasswordWithPolicy) (int, int) {
	valid := 0
	invalid := 0

	for _, p := range passwords {
		a := string(p.Password[p.Min-1])
		b := string(p.Password[p.Max-1])

		if (a == p.Sequence && b != p.Sequence) || (a != p.Sequence && b == p.Sequence) {
			valid++
		} else {
			invalid++
		}
	}

	return valid, invalid
}

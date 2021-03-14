package day6

import (
	"io"
	"os"
	"sort"
	"strings"
)

func ReadInputFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ReadInput(file)
}

func ReadInput(r io.Reader) ([]string, error) {
	result := []string{}

	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	content := string(b)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		result = append(result, line)
	}

	return result, nil
}

func ExtractGroups(lines []string) [][]string {
	result := [][]string{}

	group := []string{}
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" {
			result = append(result, group)
			group = []string{}
			continue
		}

		group = append(group, line)
	}
	if len(group) > 0 {
		result = append(result, group)
	}

	return result
}

func UniqueAnswersOfGroup(answers []string) string {
	ans := map[rune]interface{}{}

	for _, answer := range answers {
		for _, r := range answer {
			ans[r] = nil
		}
	}

	ansList := []string{}
	for r := range ans {
		ansList = append(ansList, string(r))
	}

	sort.Strings(ansList)

	return strings.Join(ansList, "")
}

package day06

import (
	"strings"
)

const (
	DEBUG1 bool = false
	DEBUG2 bool = false
)

type Range struct {
	Min int
	Max int
}

func Part1(input string) int {
	lines := parse(input)
	_ = lines
	return 0
}

func Part2(input string) int {
	lines := parse(input)
	_ = lines
	return 0
}

func parse(input string) []string {
	input = strings.TrimSpace(input)

	return strings.Split(input, "\n")
}

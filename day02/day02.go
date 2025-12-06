package day02

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	DEBUG1 bool = false
	DEBUG2 bool = false
)

func Part1(input string) int {
	entries := parse(input)
	sum := 0

	for _, entry := range entries {
		// Parse values
		values := strings.Split(entry, "-")
		start, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatalf("Error converting: %v", err)
			return 0
		}
		end, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatalf("Error converting: %v", err)
			return 0
		}

		// Checking entries
		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			if len(str)%2 != 0 {
				continue
			}

			if str[len(str)/2:] == str[:len(str)/2] {
				if DEBUG1 {
					fmt.Println(str)
				}
				sum += i
			}
		}
	}

	return sum
}

func Part2(input string) int {
	entries := parse(input)
	sum := 0

	for _, entry := range entries {
		// Parse values
		values := strings.Split(entry, "-")
		start, err := strconv.Atoi(values[0])
		if err != nil {
			log.Fatalf("Error converting: %v", err)
			return 0
		}
		end, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatalf("Error converting: %v", err)
			return 0
		}

		// Checking entries
		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			length := len(str)
			half := length / 2

			for j := 1; j <= half; j++ {
				if length%j != 0 {
					continue
				}

				repeat := length / j

				pattern := str[:j]
				testStr := strings.Repeat(pattern, repeat)

				if str == testStr {
					sum += i
					break
				}
			}
		}
	}

	return sum
}

func parse(input string) []string {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	return strings.Split(input, ",")
}

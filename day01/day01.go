package day01

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Direction int

const (
	LEFT Direction = iota
	RIGHT
)

const (
	DEBUG1 bool = false
	DEBUG2 bool = false
)

func Part1(input string) int {
	lines := parse(input)

	dial := 50
	counter := 0
	for _, line := range lines {
		fChar := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Error converting num: %v", err)
		}

		if fChar == 'L' {
			moveDial(LEFT, value, &dial)
		} else {
			moveDial(RIGHT, value, &dial)
		}

		if dial == 0 {
			counter++
		}

		if DEBUG1 {
			fmt.Println("First Char: ", string(fChar))
			fmt.Println("Dial: ", dial)
			fmt.Println("Value: ", value)
			fmt.Println("Counter: ", counter)
		}
	}

	return counter
}

func Part2(input string) int {
	lines := parse(input)
	dial := 50
	counter := 0
	var direction Direction
	for _, line := range lines {
		fChar := line[0]
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Error converting num: %v", err)
		}

		if fChar == 'L' {
			direction = LEFT
		} else {
			direction = RIGHT
		}

		counter += moveAndCheck(direction, value, &dial)

		if DEBUG2 {
			fmt.Println("First Char: ", string(fChar))
			fmt.Println("Dial: ", dial)
			fmt.Println("Value: ", value)
			fmt.Println("Counter: ", counter)
		}
	}

	return counter
}

func moveAndCheck(direction Direction, value int, dial *int) int {
	start := *dial
	var first int
	if direction == LEFT {
		if start == 0 {
			first = 100
		} else {
			first = start
		}
	} else { // RIGHT
		if start == 0 {
			first = 100
		} else {
			first = 100 - start
		}
	}

	counter := 0
	if value >= first {
		counter = 1 + (value-first)/100
	}

	moveDial(direction, value, dial)

	return counter
}

func moveDial(direction Direction, value int, dial *int) {
	if direction == LEFT {
		*dial = mod(*dial-value, 100)
	} else {
		*dial = mod(*dial+value, 100)
	}
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func parse(input string) []string {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	return strings.Split(input, "\n")
}

package day04

import (
	"strings"
)

const (
	DEBUG1 bool = false
	DEBUG2 bool = false
)

type Item rune

const (
	ROLL  Item = '@'
	EMPTY Item = '.'
)

type Location struct {
	X int
	Y int
}

var neighbors = [8]Location{{0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}, {-1, 0}, {-1, 1}}

func Part1(input string) int {
	space := parse(input)

	total := 0
	for i := range space {
		for j := range space[i] {
			if space[i][j] != ROLL {
				continue
			}

			n := countNeighbors(space, Location{j, i})
			if n < 4 {
				total++
			}
		}
	}

	return total
}

func Part2(input string) int {
	space := parse(input)

	total := 0
	current := -1

	for current != 0 {
		current = 0
		validLocs := []Location{}
		for i := range space {
			for j := range space[i] {
				if space[i][j] != ROLL {
					continue
				}

				n := countNeighbors(space, Location{j, i})
				if n < 4 {
					validLocs = append(validLocs, Location{j, i})
					current++
				}
			}
		}

		total += current
		for _, l := range validLocs {
			space[l.Y][l.X] = EMPTY
		}
	}

	return total
}

func parse(input string) [][]Item {
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	lines := strings.Split(input, "\n")

	space := make([][]Item, len(lines))
	for i, line := range lines {
		rs := []rune(line)
		space[i] = make([]Item, len(rs))
		for j, c := range rs {
			space[i][j] = Item(c)
		}
	}

	return space
}

func countNeighbors(space [][]Item, l Location) int {
	width := len(space[0])
	height := len(space)
	count := 0

	for _, n := range neighbors {
		if !inBounds(width, height, l, n) {
			continue
		}

		nx := l.X + n.X
		ny := l.Y + n.Y

		if space[ny][nx] == ROLL {
			count++
		}
	}

	return count
}

func inBounds(width, height int, l, n Location) bool {
	nx := l.X + n.X
	ny := l.Y + n.Y

	return nx >= 0 && nx < width && ny >= 0 && ny < height
}

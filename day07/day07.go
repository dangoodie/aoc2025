package day07

import (
	"fmt"
	"strings"
)

const (
	DEBUG1 bool = false
	DEBUG2 bool = false
)

type Location struct {
	X int
	Y int
}

const (
	START rune = 'S'
	SPLIT rune = '^'
	BEAM  rune = '|'
	EMPTY rune = '.'
)

func Part1(input string) int {
	lines := parse(input)
	mtx := makeRuneMatrix(lines)
	drawBeams(mtx)

	return countSplits(mtx)
}

func Part2(input string) int {
	lines := parse(input)
	mtx := makeRuneMatrix(lines)
	return countTimelines(mtx)
}

func parse(input string) []string {
	input = strings.TrimSpace(input)

	return strings.Split(input, "\n")
}

func printMatrix(mtx [][]rune) {
	for _, row := range mtx {
		fmt.Println(string(row))
	}
}

func makeRuneMatrix(lines []string) [][]rune {
	mtx := make([][]rune, len(lines))
	for i, line := range lines {
		mtx[i] = []rune(line)
	}
	return mtx
}

func lookDown(lines []string, loc Location) rune {
	height := len(lines)
	if loc.Y+1 >= height {
		return EMPTY
	}
	return rune(lines[loc.Y+1][loc.X])
}

func splitBeam(mtx [][]rune, loc Location) {
	if mtx[loc.Y][loc.X] != SPLIT {
		if DEBUG1 || DEBUG2 {
			fmt.Println("Fail splitBeam: ", loc)
		}
		return
	}

	if mtx[loc.Y][loc.X-1] == EMPTY {
		mtx[loc.Y][loc.X-1] = BEAM
	}
	if mtx[loc.Y][loc.X+1] == EMPTY {
		mtx[loc.Y][loc.X+1] = BEAM
	}
}

func findStart(mtx [][]rune) (Location, bool) {
	for y, row := range mtx {
		for x, cell := range row {
			if cell == START {
				return Location{X: x, Y: y}, true
			}
		}
	}

	return Location{}, false
}

func drawBeams(mtx [][]rune) {
	height := len(mtx)
	if height == 0 {
		return
	}
	width := len(mtx[0])

	start, ok := findStart(mtx)
	if !ok {
		if DEBUG1 || DEBUG2 {
			fmt.Println("No START found")
		}
		return
	}

	beams := map[int]bool{
		start.X: true,
	}

	for y := start.Y + 1; y < height && len(beams) > 0; y++ {
		nextBeams := make(map[int]bool)

		for x := range beams {
			if x < 0 || x >= width {
				continue
			}

			cell := mtx[y][x]

			if cell == EMPTY {
				mtx[y][x] = BEAM
			}

			if cell == SPLIT {
				splitBeam(mtx, Location{X: x, Y: y})

				if x-1 >= 0 {
					nextBeams[x-1] = true
				}
				if x+1 < width {
					nextBeams[x+1] = true
				}

			} else {
				nextBeams[x] = true
			}
		}
		beams = nextBeams
	}

	if DEBUG1 {
		fmt.Println("After drawing beams:")
		printMatrix(mtx)
	}

}

func countSplits(mtx [][]rune) int {
	height := len(mtx)
	if height == 0 {
		return 0
	}
	width := len(mtx[0])

	splits := 0

	for y := 1; y < height; y++ {
		for x := 0; x < width; x++ {
			if mtx[y][x] != SPLIT {
				continue
			}

			above := mtx[y-1][x]
			if above == BEAM || above == START {
				splits++
			}
		}
	}

	return splits
}

func countTimelines(mtx [][]rune) int {
	height := len(mtx)
	if height == 0 {
		return 0
	}
	width := len(mtx[0])

	start, ok := findStart(mtx)
	if !ok {
		if DEBUG1 || DEBUG2 {
			fmt.Println("No START found")
		}
		return 0
	}

	// timelines[x] = number of timelines at column x on the current row
	timelines := map[int]int{
		start.X: 1, // 1 timeline starting under S
	}

	exitTotal := 0

	// Walk from the row below S down through all rows
	for y := start.Y + 1; y < height; y++ {
		next := make(map[int]int)

		for x, count := range timelines {
			if count == 0 {
				continue
			}

			// If somehow we got out of bounds horizontally, those timelines have already exited
			if x < 0 || x >= width {
				exitTotal += count
				continue
			}

			cell := mtx[y][x]

			if cell == SPLIT {
				// Split left
				leftX := x - 1
				if leftX >= 0 {
					next[leftX] += count
				} else {
					// Left branch exits manifold
					exitTotal += count
				}

				// Split right
				rightX := x + 1
				if rightX < width {
					next[rightX] += count
				} else {
					// Right branch exits manifold
					exitTotal += count
				}
			} else {
				// Just continue straight down
				next[x] += count
			}
		}

		timelines = next
	}

	// After processing the last row, any remaining timelines step off the bottom.
	for _, count := range timelines {
		exitTotal += count
	}

	return exitTotal
}

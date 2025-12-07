package day06

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	DEBUG1 bool = false
	DEBUG2 bool = false
)

type Problem rune

const (
	ADD Problem = '+'
	MUL Problem = '*'
)

func Part1(input string) int {
	lines := parse(input)
	problems := parseSymbols(lines[len(lines)-1])
	numsMtx := parseNums(lines[:len(lines)-1])

	if DEBUG1 {
		for _, row := range numsMtx {
			for _, num := range row {
				fmt.Printf("%v ", num)
			}
			fmt.Println("")
		}

		for _, p := range problems {
			fmt.Printf("%c ", p)
		}
		fmt.Println("")
	}

	total := 0
	for i, p := range problems {
		total += computeCol(numsMtx, i, p)
	}

	return total
}

func Part2(input string) int {
	lines := parse2(input)
	mtx := parseNums2(lines)
	probLine := strings.TrimSpace(lines[len(lines)-2])
	problems := parseSymbols(probLine)

	if DEBUG2 {
		fmt.Println("Matrix: ", mtx)
		fmt.Println("Problems: ", problems)
	}

	total := 0
	for i, p := range problems {
		if p == ADD {
			sum := 0
			for _, num := range mtx[i] {
				sum += num
			}
			total += sum
		}

		if p == MUL {
			prod := 1
			for _, num := range mtx[i] {
				prod *= num
			}
			total += prod
		}
	}

	return total
}

func parse(input string) []string {
	input = strings.TrimSpace(input)

	return strings.Split(input, "\n")
}

func parse2(input string) []string {
	return strings.Split(input, "\n")
}

func parseSymbols(line string) []Problem {
	symbols := strings.Fields(line)
	var problems []Problem

	for _, s := range symbols {
		problems = append(problems, Problem(s[0]))
	}

	return problems
}

func parseNums(lines []string) [][]int {
	var matrix [][]int

	for _, line := range lines {
		arr := strings.Fields(line)
		var row []int
		for _, s := range arr {
			num, err := strconv.Atoi(s)
			if err != nil {
				panic(1)
			}
			row = append(row, num)
		}
		matrix = append(matrix, row)
	}

	return matrix
}

func computeCol(mtx [][]int, col int, p Problem) int {
	total := 0

	for _, row := range mtx {
		// handle if 0
		if total == 0 {
			total += row[col]
			continue
		}

		if p == ADD {
			total += row[col]
		} else {
			total *= row[col]
		}
	}

	return total
}
func parseNums2(lines []string) [][]int {
	var mtx [][]int
	width := len(lines[0])
	height := len(lines)

	// down a column byte by byte, constructing a slice of bytes
	// if End of Rows hit, construct num and append to row
	// if Symbol hit, construct num, append to row, and prepend row to mtx

	var row []int
	found := false

	for col := width - 1; col >= 0; col-- {
		var digits []byte
		for j := range height - 1 {
			r := lines[j][col]
			if r == ' ' {
				continue
			}

			if unicode.IsDigit(rune(r)) {
				digits = append(digits, r)
			}

			if r == byte(ADD) || r == byte(MUL) { // End of Row found
				found = true
			}
		}

		if len(digits) == 0 {
			continue
		}

		//flush digits into number
		num, err := strconv.Atoi(string(digits))
		if err != nil {
			fmt.Println("Error converting to digits")
			fmt.Println(digits)
			panic(1)
		}

		//append to row
		row = append(row, num)

		// if found prepend to matrix
		if found {
			mtx = append([][]int{row}, mtx...)
			row = []int{}
			found = false
		}
	}

	return mtx
}

package main

import (
	"fmt"
	"log"

	"github.com/dangoodie/aoc2025/day04"
	"github.com/dangoodie/aoc2025/internal/input"
)

func main() {
	data, err := input.Read("day04/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	fmt.Println("Day 04")
	fmt.Println("Part 1:", day04.Part1(data))
	fmt.Println("Part 2:", day04.Part2(data))
}

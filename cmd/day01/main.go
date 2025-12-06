package main

import (
	"fmt"
	"log"

	"github.com/dangoodie/aoc2025/day01"
	"github.com/dangoodie/aoc2025/internal/input"
)

func main() {
	data, err := input.Read("day01/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	fmt.Println("Day 01")
	fmt.Println("Part 1:", day01.Part1(data))
	fmt.Println("Part 2:", day01.Part2(data))
}

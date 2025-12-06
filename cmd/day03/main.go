package main

import (
	"fmt"
	"log"

	"github.com/dangoodie/aoc2025/day03"
	"github.com/dangoodie/aoc2025/internal/input"
)

func main() {
	data, err := input.Read("day03/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	fmt.Println("Day 03")
	fmt.Println("Part 1:", day03.Part1(data))
	fmt.Println("Part 2:", day03.Part2(data))
}

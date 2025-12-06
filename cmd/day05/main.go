package main

import (
	"fmt"
	"log"

	"github.com/dangoodie/aoc2025/day05"
	"github.com/dangoodie/aoc2025/internal/input"
)

func main() {
	data, err := input.Read("day05/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	fmt.Println("Day 05")
	fmt.Println("Part 1:", day05.Part1(data))
	fmt.Println("Part 2:", day05.Part2(data))
}

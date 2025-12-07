package main

import (
	"fmt"
	"log"

	"github.com/dangoodie/aoc2025/day06"
	"github.com/dangoodie/aoc2025/internal/input"
)

func main() {
	data, err := input.Read("day06/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	fmt.Println("Day 06")
	fmt.Println("Part 1:", day06.Part1(data))
	fmt.Println("Part 2:", day06.Part2(data))
}

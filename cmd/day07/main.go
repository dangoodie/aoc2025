package main

import (
	"fmt"
	"log"

	"github.com/dangoodie/aoc2025/day07"
	"github.com/dangoodie/aoc2025/internal/input"
)

func main() {
	data, err := input.Read("day07/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	fmt.Println("Day 07")
	fmt.Println("Part 1:", day07.Part1(data))
	fmt.Println("Part 2:", day07.Part2(data))
}

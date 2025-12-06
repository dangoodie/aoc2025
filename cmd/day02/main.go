package main

import (
	"fmt"
	"log"

	"github.com/dangoodie/aoc2025/day02"
	"github.com/dangoodie/aoc2025/internal/input"
)

func main() {
	data, err := input.Read("day02/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	fmt.Println("Day 02")
	fmt.Println("Part 1:", day02.Part1(data))
	fmt.Println("Part 2:", day02.Part2(data))
}

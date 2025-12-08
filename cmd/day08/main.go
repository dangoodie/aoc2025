package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dangoodie/aoc2025/day08"
	"github.com/dangoodie/aoc2025/internal/input"
)

func main() {
	data, err := input.Read("day08/input.txt")
	if err != nil {
		log.Fatalf("read input: %v", err)
	}

	fmt.Println("Day 08")
	start := time.Now()
	fmt.Println("Part 1:", day08.Part1(data))
	part1Time := time.Since(start)
	start = time.Now()
	fmt.Println("Part 2:", day08.Part2(data))
	part2Time := time.Since(start)
	fmt.Println("-----------Stats-----------")
	fmt.Println("Part 1 Time: ", part1Time)
	fmt.Println("Part 2 Time: ", part2Time)
	fmt.Println("Total: ", part1Time+part2Time)
}

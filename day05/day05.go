package day05

import (
	"sort"
	"strconv"
	"strings"
)

const (
	DEBUG1 bool = false
	DEBUG2 bool = false
)

type Range struct {
	Min int
	Max int
}

func Part1(input string) int {
	lines := parse(input)
	ranges := parseRanges(lines)
	ids := parseIDs(lines)

	total := 0

	for _, id := range ids {
		for _, r := range ranges {
			if inRange(id, r) {
				total++
				break
			}
		}

	}
	return total
}

func Part2(input string) int {
	lines := parse(input)
	ranges := parseRanges(lines)

	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].Min == ranges[j].Min {
			return ranges[i].Max < ranges[j].Max
		}

		return ranges[i].Min < ranges[j].Min
	})

	total := 0
	cur := ranges[0]

	for _, r := range ranges[1:] {
		if r.Min <= cur.Max+1 {
			if r.Max > cur.Max {
				cur.Max = r.Max
			}
			continue
		}
		total += cur.Max - cur.Min + 1
		cur = r
	}

	total += cur.Max - cur.Min + 1
	return total
}

func parse(input string) []string {
	input = strings.TrimSpace(input)

	return strings.Split(input, "\n")
}

func parseRanges(lines []string) []Range {
	var ranges []Range
	for _, line := range lines {
		if line == "" {
			break
		}

		n := strings.Split(line, "-")
		min, err := strconv.Atoi(n[0])
		if err != nil {
			panic(1)
		}
		max, err := strconv.Atoi(n[1])
		if err != nil {
			panic(1)
		}
		r := Range{
			Min: min,
			Max: max,
		}

		ranges = append(ranges, r)
	}

	return ranges
}

func parseIDs(lines []string) []int {
	idx := 0
	for i, line := range lines {
		if line == "" {
			idx = i
			break
		}
	}

	var ids []int
	for _, line := range lines[idx+1:] {
		id, err := strconv.Atoi(line)
		if err != nil {
			panic(1)
		}
		ids = append(ids, id)
	}

	return ids
}

func inRange(id int, r Range) bool {
	return r.Min <= id && id <= r.Max
}

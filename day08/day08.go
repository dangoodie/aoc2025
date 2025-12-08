package day08

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/dangoodie/aoc2025/internal/dsu"
)

const (
	DEBUG1 bool = false
	DEBUG2 bool = false
)

type Point struct {
	X int
	Y int
	Z int
}

type Connection struct {
	D2 int
	A  int //index
	B  int //index
}

func Part1(input string) int {
	points := parse(input)
	var conns []Connection

	// create connections
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			d2 := calcDist(points[i], points[j])
			conns = append(conns, Connection{
				D2: d2,
				A:  i,
				B:  j,
			})
		}
	}

	// sort connections by shortest distance
	sort.Slice(conns, func(i, j int) bool {
		return conns[i].D2 < conns[j].D2
	})

	// process connections
	dsu := dsu.NewDSU(len(points))
	for i := 0; i < 1000; i++ {
		conn := conns[i]
		dsu.Union(conn.A, conn.B)
	}

	// aggregate sizes
	sizesMap := make(map[int]int)
	for i := range points {
		root := dsu.Find(i)
		sizesMap[root] = dsu.Size(root)
	}

	var sizes []int
	for _, s := range sizesMap {
		sizes = append(sizes, s)
	}
	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	return sizes[0] * sizes[1] * sizes[2]
}

func Part2(input string) int {
	points := parse(input)
	var conns []Connection

	// create connections
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			d2 := calcDist(points[i], points[j])
			conns = append(conns, Connection{
				D2: d2,
				A:  i,
				B:  j,
			})
		}
	}

	// sort connections by shortest distance
	sort.Slice(conns, func(i, j int) bool {
		return conns[i].D2 < conns[j].D2
	})

	dsu := dsu.NewDSU(len(points))
	junctions := len(points)
	var lastA, lastB int

	for _, c := range conns {
		merged := dsu.Union(c.A, c.B)
		if merged {
			junctions--
			if junctions == 1 {
				lastA, lastB = c.A, c.B
				break
			}
		}
	}

	return points[lastA].X * points[lastB].X
}

func calcDist(a, b Point) int {
	dx := a.X - b.X
	dy := a.Y - b.Y
	dz := a.Z - b.Z

	return dx*dx + dy*dy + dz*dz
}

func parse(input string) []Point {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	var points []Point
	for _, line := range lines {
		vals := strings.Split(line, ",")
		p := convertPoint(vals)
		points = append(points, p)
	}

	return points
}

func convertPoint(vals []string) Point {
	if len(vals) != 3 {
		fmt.Println("Error parsing val")
		panic(1)
	}
	x, err := strconv.Atoi(vals[0])
	if err != nil {
		fmt.Println("Error parsing val")
		panic(1)
	}
	y, err := strconv.Atoi(vals[1])
	if err != nil {
		fmt.Println("Error parsing val")
		panic(1)
	}
	z, err := strconv.Atoi(vals[2])
	if err != nil {
		fmt.Println("Error parsing val")
		panic(1)
	}

	return Point{X: x, Y: y, Z: z}
}

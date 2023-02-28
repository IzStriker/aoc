package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type coordinate struct {
	x int
	y int
}

type pair struct {
	s coordinate
	b coordinate
}

type line struct {
	c1 coordinate
	c2 coordinate
}

func main() {
	pairs, beacons := parse()
	part1(pairs, beacons)
	part2(pairs, beacons)

}

func part1(pairs []pair, beacons map[coordinate]bool) {
	row := 2000000
	positions := make(map[coordinate]bool)

	for _, value := range pairs {
		distance := getDistance(value.b, value.s)
		if getDistance(value.s, coordinate{value.s.x, row}) <= distance {
			if _, exists := beacons[coordinate{value.s.x, row}]; !exists {
				positions[coordinate{value.s.x, row}] = true
			}
			x := value.s.x - 1

			// expand left
			for getDistance(value.s, coordinate{x, row}) <= distance {
				if _, exists := beacons[coordinate{x, row}]; !exists {
					positions[coordinate{x, row}] = true
				}
				x -= 1
			}

			x = value.s.x + 1
			// expand right
			for getDistance(value.s, coordinate{x, row}) <= distance {
				if _, exists := beacons[coordinate{x, row}]; !exists {
					positions[coordinate{x, row}] = true
				}
				x += 1
			}
		}
	}

	fmt.Println(len(positions))
}

func part2(pairs []pair, beacons map[coordinate]bool) {
	// https://www.reddit.com/r/adventofcode/comments/zmcn64/comment/j0sprzi/?utm_source=share&utm_medium=web2x&context=3

	// get the perimeter one cell out of each sphere, calculate the liens that make up the perimeter
	// calculate the intersection of each line with every other line
	// if the intersection is within the search space, and not closer to a sensor than the sensors closest beacon, then it's a valid point
	// there should only be one valid point

	searchSpace := 4_000_000
	lines := make([]line, 0)
	for _, pair := range pairs {
		radius := getDistance(pair.s, pair.b) + 1

		top := coordinate{x: pair.s.x, y: pair.s.y - radius}
		bottom := coordinate{x: pair.s.x, y: pair.s.y + radius}
		left := coordinate{x: pair.s.x - radius, y: pair.s.y}
		right := coordinate{x: pair.s.x + radius, y: pair.s.y}
		// top left
		lines = append(lines, line{left, top})
		// top right
		lines = append(lines, line{top, right})
		// bottom right
		lines = append(lines, line{right, bottom})
		// bottom left
		lines = append(lines, line{bottom, left})

	}

	// calculate intersection
	var intersections []coordinate
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			intersects, coord := getIntersection(lines[i], lines[j])
			if i != j &&
				intersects &&
				coord.x >= 0 &&
				coord.y >= 0 &&
				coord.x <= searchSpace &&
				coord.y <= searchSpace {
				intersections = append(intersections, coord)
			}
		}
	}

	intersectionMap := make(map[coordinate]bool)

	for i := 0; i < len(intersections); i++ {

		point := intersections[i]
		// check that point isn't closer to a sensor than the sensors closest beacon
		// if it is, then it's not a valid point

		fail := false
		for _, pair := range pairs {
			radius := getDistance(pair.s, pair.b)
			distance := getDistance(pair.s, point)

			if distance <= radius {
				fail = true
			}
		}

		// // check if top, bottom, left, right are in the list of intersections
		if fail {
			continue
		}

		intersectionMap[point] = true
	}

	for point, _ := range intersectionMap {
		fmt.Println("(", point.x, ",", point.y, ")")
		fmt.Println(point.x*4_000_000 + point.y)
	}

}
func getIntersection(l1 line, l2 line) (bool, coordinate) {
	// https://en.wikipedia.org/wiki/Line%E2%80%93line_intersection#Given_two_points_on_each_line

	// calculate denominator
	l1xDiff := l1.c1.x - l1.c2.x // x1 - x2
	l2xDiff := l2.c1.x - l2.c2.x // x3 - x4
	l1yDiff := l1.c1.y - l1.c2.y // y1 - y2
	l2yDiff := l2.c1.y - l2.c2.y // y3 - y4

	denom := calcDeterminant(l1xDiff, l1yDiff, l2xDiff, l2yDiff)
	if denom == 0 {
		return false, coordinate{}
	}

	l1Det := calcDeterminant(l1.c1.x, l1.c1.y, l1.c2.x, l1.c2.y)
	l2Det := calcDeterminant(l2.c1.x, l2.c1.y, l2.c2.x, l2.c2.y)

	return true, coordinate{
		x: calcDeterminant(l1Det, l1xDiff, l2Det, l2xDiff) / denom,
		y: calcDeterminant(l1Det, l1yDiff, l2Det, l2yDiff) / denom,
	}
}

func calcDeterminant(a, b, c, d int) int {
	// |a b|
	// |c d|
	return (a * d) - (b * c)
}

func getDistance(a coordinate, b coordinate) int {
	x1 := max(a.x, b.x)
	x2 := min(a.x, b.x)
	y1 := max(a.y, b.y)
	y2 := min(a.y, b.y)

	return (x1 - x2) + (y1 - y2)
}

func min(a int, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a int, b int) int {
	if b > a {
		return b
	}
	return a
}

func parse() ([]pair, map[coordinate]bool) {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r, _ := regexp.Compile(`-?\d+`)

	pairs := make([]pair, 0)
	beacons := make(map[coordinate]bool)
	for scanner.Scan() {
		coords := toInt(r.FindAllString(scanner.Text(), 4))
		p := pair{
			s: coordinate{x: coords[0], y: coords[1]},
			b: coordinate{x: coords[2], y: coords[3]},
		}
		beacons[coordinate{x: coords[2], y: coords[3]}] = true
		pairs = append(pairs, p)
	}
	return pairs, beacons
}

func toInt(inputs []string) []int {
	outputs := make([]int, 0)
	for _, s := range inputs {
		value, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		outputs = append(outputs, value)
	}
	return outputs
}

func toStrings(inputs []int) []string {
	outputs := make([]string, 0)
	for _, i := range inputs {
		outputs = append(outputs, strconv.Itoa(i))
	}
	return outputs
}

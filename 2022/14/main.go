package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func main() {
	part2()
}

func min(num1 int, num2 int) int {
	if num2 < num1 {
		return num2
	}
	return num1
}

func max(num1 int, num2 int) int {
	if num2 > num1 {
		return num2
	}
	return num1
}

func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cave := make(map[coordinate]string)

	maxY := math.MinInt
	// Get Paths
	for scanner.Scan() {
		path := make([]coordinate, 0)
		for _, v := range strings.Split(scanner.Text(), " -> ") {
			point := strings.Split(v, ",")

			co := coordinate{}
			if x, err := strconv.Atoi(point[0]); err == nil {
				co.x = x
			} else {
				panic(err)
			}
			if y, err := strconv.Atoi(point[1]); err == nil {
				co.y = y
				if y > maxY {
					maxY = y
				}
			} else {
				panic(err)
			}
			path = append(path, co)
		}

		for i := 0; i < len(path)-1; i++ {
			if path[i].x == path[i+1].x {
				start := min(path[i].y, path[i+1].y)
				end := max(path[i].y, path[i+1].y)

				for j := start; j < end+1; j++ {
					cave[coordinate{path[i].x, j}] = "path"
				}
			} else if path[i].y == path[i+1].y {
				start := min(path[i].x, path[i+1].x)
				end := max(path[i].x, path[i+1].x)

				for j := start; j < end+1; j++ {
					cave[coordinate{j, path[i].y}] = "path"
				}
			}
		}
	}

	// Get Sand
	i := 0
	inAbyss := false
	for !inAbyss {
		atRest := false
		sand := coordinate{500, 0}

		for !atRest {
			if _, exists := cave[coordinate{sand.x, sand.y + 1}]; !exists {
				//down
				sand.y += 1
			} else if _, exists := cave[coordinate{sand.x - 1, sand.y + 1}]; !exists {
				// left
				sand.y += 1
				sand.x -= 1
			} else if _, exists := cave[coordinate{sand.x + 1, sand.y + 1}]; !exists {
				// right
				sand.y += 1
				sand.x += 1
			} else {
				atRest = true
			}

			if sand.y > maxY {
				fmt.Println(i)
				inAbyss = true
				break
			}
		}
		if inAbyss {
			break
		}

		i += 1
		cave[sand] = "sand"
	}
}
func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cave := make(map[coordinate]string)

	maxY := math.MinInt
	// Get Paths
	for scanner.Scan() {
		path := make([]coordinate, 0)
		for _, v := range strings.Split(scanner.Text(), " -> ") {
			point := strings.Split(v, ",")

			co := coordinate{}
			if x, err := strconv.Atoi(point[0]); err == nil {
				co.x = x
			} else {
				panic(err)
			}
			if y, err := strconv.Atoi(point[1]); err == nil {
				co.y = y
				if y > maxY {
					maxY = y
				}
			} else {
				panic(err)
			}
			path = append(path, co)
		}

		for i := 0; i < len(path)-1; i++ {
			if path[i].x == path[i+1].x {
				start := min(path[i].y, path[i+1].y)
				end := max(path[i].y, path[i+1].y)

				for j := start; j < end+1; j++ {
					cave[coordinate{path[i].x, j}] = "path"
				}
			} else if path[i].y == path[i+1].y {
				start := min(path[i].x, path[i+1].x)
				end := max(path[i].x, path[i+1].x)

				for j := start; j < end+1; j++ {
					cave[coordinate{j, path[i].y}] = "path"
				}
			}
		}
	}

	// Get Sand
	i := 0
	inAbyss := false
	for !inAbyss {
		atRest := false
		sand := coordinate{500, 0}

		if _, exists := cave[sand]; exists {
			fmt.Println(i)
			break
		}

		for !atRest {
			if _, exists := cave[coordinate{sand.x, sand.y + 1}]; !exists && sand.y+1 < maxY+2 {
				//down
				sand.y += 1
			} else if _, exists := cave[coordinate{sand.x - 1, sand.y + 1}]; !exists && sand.y+1 < maxY+2 {
				// left
				sand.y += 1
				sand.x -= 1
			} else if _, exists := cave[coordinate{sand.x + 1, sand.y + 1}]; !exists && sand.y+1 < maxY+2 {
				// right
				sand.y += 1
				sand.x += 1
			} else {
				atRest = true
			}

		}

		i += 1
		cave[sand] = "sand"
	}
}

func drawExample(cave map[coordinate]string) {
	for y := 0; y < 10; y++ {
		for x := 494; x < 504; x++ {
			block, exists := cave[coordinate{x, y}]
			if !exists {
				fmt.Print(".")
			} else if block == "path" {
				fmt.Print("#")
			} else if block == "sand" {
				fmt.Print("o")
			}

		}
		fmt.Println()
	}
}

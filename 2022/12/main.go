package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// heightMap := parse()
	// part1(heightMap)
	heightMap := parse()
	part2(heightMap)
}

func parse() [][]rune {
	heightMap := make([][]rune, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := make([]rune, 0)
		for _, value := range scanner.Text() {
			line = append(line, value)
		}
		heightMap = append(heightMap, line)
	}
	return heightMap
}

type coordinate struct {
	x int
	y int
}

func part1(heightMap [][]rune) {

	// find S
	start := coordinate{}
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if heightMap[i][j] == 'S' {
				start.x = j
				start.y = i
				break
			}
		}
	}

	// find E
	end := coordinate{}
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if heightMap[i][j] == 'E' {
				end.x = j
				end.y = i
				break
			}
		}
	}

	costs := make(map[coordinate]int)
	visited := make(map[coordinate]bool)
	queue := make([]coordinate, 0)
	current := coordinate{}
	current.x, current.y = start.x, start.y

	heightMap[start.y][start.x] = 'a'
	heightMap[end.y][end.x] = 'z'
	costs[start] = 0
	for !(current.x == end.x && current.y == end.y) {
		neighbours := make([]coordinate, 0)
		visited[current] = true

		// up
		if current.y-1 >= 0 && heightMap[current.y-1][current.x] <= heightMap[current.y][current.x]+1 {
			neighbours = append(neighbours, coordinate{current.x, current.y - 1})
		}

		// down
		if current.y+1 < len(heightMap) && heightMap[current.y+1][current.x] <= heightMap[current.y][current.x]+1 {
			neighbours = append(neighbours, coordinate{current.x, current.y + 1})
		}

		// left
		if current.x-1 >= 0 && heightMap[current.y][current.x-1] <= heightMap[current.y][current.x]+1 {
			neighbours = append(neighbours, coordinate{current.x - 1, current.y})
		}

		// right
		if current.x+1 < len(heightMap[current.y]) && heightMap[current.y][current.x+1] <= heightMap[current.y][current.x]+1 {
			neighbours = append(neighbours, coordinate{current.x + 1, current.y})
		}

		for _, v := range neighbours {
			cost, exists := costs[v]
			if !exists {
				costs[v] = costs[current] + 1
			} else if cost > costs[current]+1 {
				costs[v] = costs[current] + 1
			}

			inside := false
			for i := 0; i < len(queue); i++ {
				if queue[i] == v {
					inside = true
					break
				}
			}

			if !visited[v] && !inside {
				queue = append(queue, v)
			}
		}

		// fmt.Println("current", current)
		// fmt.Println("neighbours", neighbours)
		// fmt.Println("queue", queue)
		// fmt.Println("costs", costs)
		// bufio.NewReader(os.Stdin).ReadBytes('\n')

		current, queue = queue[0], queue[1:]
	}

	fmt.Println(costs[end])
}
func part2(heightMap [][]rune) {

	// find all a
	starts := make([]coordinate, 0)
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if heightMap[i][j] == 'S' || heightMap[i][j] == 'a' {
				starts = append(starts, coordinate{j, i})
				if heightMap[i][j] == 'S' {
					heightMap[i][j] = 'a'
				}
			}
		}
	}

	// find E
	end := coordinate{}
	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if heightMap[i][j] == 'E' {
				end.x = j
				end.y = i
				break
			}
		}
	}

	min := 100_000_000_00
	heightMap[end.y][end.x] = 'z'

	for _, start := range starts {
		costs := make(map[coordinate]int, 0)
		visited := make(map[coordinate]bool, 0)
		queue := make([]coordinate, 0)

		current := coordinate{start.x, start.y}
		costs[start] = 0
		for !(current.x == end.x && current.y == end.y) {
			neighbours := make([]coordinate, 0)
			visited[current] = true

			// up
			if current.y-1 >= 0 && heightMap[current.y-1][current.x] <= heightMap[current.y][current.x]+1 {
				neighbours = append(neighbours, coordinate{current.x, current.y - 1})
			}

			// down
			if current.y+1 < len(heightMap) && heightMap[current.y+1][current.x] <= heightMap[current.y][current.x]+1 {
				neighbours = append(neighbours, coordinate{current.x, current.y + 1})
			}

			// left
			if current.x-1 >= 0 && heightMap[current.y][current.x-1] <= heightMap[current.y][current.x]+1 {
				neighbours = append(neighbours, coordinate{current.x - 1, current.y})
			}

			// right
			if current.x+1 < len(heightMap[current.y]) && heightMap[current.y][current.x+1] <= heightMap[current.y][current.x]+1 {
				neighbours = append(neighbours, coordinate{current.x + 1, current.y})
			}

			for _, v := range neighbours {
				cost, exists := costs[v]
				if !exists {
					costs[v] = costs[current] + 1
				} else if cost > costs[current]+1 {
					costs[v] = costs[current] + 1
				}

				inside := false
				for i := 0; i < len(queue); i++ {
					if queue[i] == v {
						inside = true
						break
					}
				}

				if !visited[v] && !inside {
					queue = append(queue, v)
				}
			}

			if len(queue) > 0 {
				current = queue[0]
				queue = queue[1:]
			} else {
				break
			}
		}

		if costs[end] < min && costs[end] != 0 {
			min = costs[end]
		}
	}
	fmt.Println(min)

}

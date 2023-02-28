package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part1() {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file), "\n")
	grid := make([][]int, 0)

	for _, line := range input {
		row := make([]int, 0)
		for _, digit := range line {
			num, err := strconv.Atoi(string(digit))
			if err != nil {
				continue
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}

	fmt.Println("size", len(grid), len(grid[0]))
	seen := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			// if outer
			if i == 0 || j == 0 || i == len(grid)-1 || j == len(grid[i])-1 {
				seen += 1
				continue
			}

			// seen from top
			row := i
			failed := false
			for row > 0 {
				row--
				if grid[i][j] <= grid[row][j] {
					failed = true
					break
				}
			}
			if !failed {
				seen += 1
				continue
			}

			// seen from bottom
			row = i
			failed = false
			for row < len(grid)-1 {
				row++
				if grid[i][j] <= grid[row][j] {
					failed = true
					break
				}
			}
			if !failed {
				seen += 1
				continue
			}

			// seen from left
			column := j
			failed = false
			for column > 0 {
				column--
				if grid[i][j] <= grid[i][column] {
					failed = true
					break
				}
			}
			if !failed {
				seen += 1
				continue
			}

			// seen from right
			column = j
			failed = false
			for column < len(grid[i])-1 {
				column++
				if grid[i][j] <= grid[i][column] {
					failed = true
					break
				}
			}
			if !failed {
				seen += 1
				continue
			}
		}
	}
	fmt.Println(seen)
}

func part2() {

	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(file), "\n")
	grid := make([][]int, 0)

	for _, line := range input {
		row := make([]int, 0)
		for _, digit := range line {
			num, err := strconv.Atoi(string(digit))
			if err != nil {
				continue
			}
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	max := -1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			top := 0
			bottom := 0
			left := 0
			right := 0

			// seen from top
			row := i
			for row > 0 {
				row--

				if grid[i][j] <= grid[row][j] {
					top++
					break
				}
				top++
			}

			// seen from bottom
			row = i
			for row < len(grid)-1 {
				row++
				if grid[i][j] <= grid[row][j] {
					bottom++
					break
				}
				bottom++
			}

			// seen from left
			column := j
			for column > 0 {
				column--
				if grid[i][j] <= grid[i][column] {
					left++
					break
				}

				left++
			}

			// seen from right
			column = j
			for column < len(grid[i])-1 {
				column++
				if grid[i][j] <= grid[i][column] {
					right++
					break
				}

				right++
			}
			score := top * bottom * left * right
			if score > max {
				max = score
			}
		}
	}
	fmt.Println(max)
}

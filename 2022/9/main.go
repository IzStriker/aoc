package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	// part1()
	part2()
}

func part1() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	visited := make(map[string]bool, 0)
	hx, hy := 0, 0
	tx, ty := 0, 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		move, err := strconv.Atoi(input[1])
		if err != nil {
			panic(err)
		}

		for i := 0; i < move; i++ {
			if input[0] == "U" {
				hy += 1
			} else if input[0] == "D" {
				hy -= 1
			} else if input[0] == "L" {
				hx -= 1
			} else if input[0] == "R" {
				hx += 1
			}

			if int(math.Abs(float64(tx-hx))) > 1 || int(math.Abs(float64(ty-hy))) > 1 {
				if hx == tx {
					if hy > ty {
						ty += 1
					} else if hy < ty {
						ty -= 1
					}
				} else if hy == ty {
					if hx > tx {
						tx += 1
					} else if hx < tx {
						tx -= 1
					}
				} else {
					if hx < tx {
						tx -= 1
					} else if hx > tx {
						tx += 1
					}

					if hy < ty {
						ty -= 1
					} else if hy > ty {
						ty += 1
					}
				}
			}
			visited[fmt.Sprint(tx, ty)] = true
			fmt.Println(input, hx, hy, tx, ty)
		}
	}
	fmt.Println(len(visited))

}

func part2() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	visited := make(map[string]bool, 0)
	x := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	y := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		move, err := strconv.Atoi(input[1])
		if err != nil {
			panic(err)
		}

		for i := 0; i < move; i++ {
			if input[0] == "U" {
				y[0] += 1
			} else if input[0] == "D" {
				y[0] -= 1
			} else if input[0] == "L" {
				x[0] -= 1
			} else if input[0] == "R" {
				x[0] += 1
			}

			for i := 1; i < len(x); i++ {

				if int(math.Abs(float64(x[i]-x[i-1]))) > 1 || int(math.Abs(float64(y[i]-y[i-1]))) > 1 {
					if x[i-1] == x[i] {
						if y[i-1] > y[i] {
							y[i] += 1
						} else if y[i-1] < y[i] {
							y[i] -= 1
						}
					} else if y[i-1] == y[i] {
						if x[i-1] > x[i] {
							x[i] += 1
						} else if x[i-1] < x[i] {
							x[i] -= 1
						}
					} else {
						if x[i-1] < x[i] {
							x[i] -= 1
						} else if x[i-1] > x[i] {
							x[i] += 1
						}

						if y[i-1] < y[i] {
							y[i] -= 1
						} else if y[i-1] > y[i] {
							y[i] += 1
						}
					}
				}
				if i == 9 {

					visited[fmt.Sprint(x[i], y[i])] = true
				}
			}
		}
	}
	fmt.Println(len(visited))

}

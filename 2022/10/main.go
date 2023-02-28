package main

import (
	"bufio"
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

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cycle := 0
	x := 1
	store := make(map[int]int, 0)
	for scanner.Scan() {
		cycle += 1

		input := strings.Split(scanner.Text(), " ")
		store[cycle] = x

		if input[0] == "addx" {
			value, err := strconv.Atoi(input[1])
			if err != nil {
				panic(err)
			}
			cycle += 1
			store[cycle] = x
			x += value
		}

	}
	fmt.Println(store[20], store[60], store[100], store[140], store[180], store[220])
	fmt.Println((20 * store[20]) + (60 * store[60]) + (100 * store[100]) + (140 * store[140]) + (180 * store[180]) + (220 * store[220]))
}

func part2() {

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cycle := 0
	x := 1
	store := make(map[int]int, 0)
	for scanner.Scan() {
		cycle += 1

		input := strings.Split(scanner.Text(), " ")
		store[cycle] = x

		if input[0] == "addx" {
			value, err := strconv.Atoi(input[1])
			if err != nil {
				panic(err)
			}
			cycle += 1
			store[cycle] = x
			x += value
		}
	}

	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			value, _ := store[(i*40)+j+1]
			if j <= value+1 && j >= value-1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

			if j == 39 {
				fmt.Println()
			}
		}
	}
}

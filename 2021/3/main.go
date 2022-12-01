package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)


	var count [12]int
	for scanner.Scan() {
		line := scanner.Text()

		for i := 0; i < 12; i++ {
			if line[i] == "1"[0] {
				count[i]++
			} else {
				count[i]--
			}
		}
	}
	
	var gamma uint
	var epsilon uint 

	for _, ele := range count {
		if ele > 0 {
			gamma = gamma << 1
			gamma += 1
		} else {
			gamma = gamma << 1
		}
	}
	epsilon = uint(0b111111111111) ^ gamma
	fmt.Print("part 1:", epsilon * gamma)
}

func part2() {
	
}
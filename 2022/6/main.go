package main

import "fmt"

func main() {
	part1()
	part2()
}

func part1() {
	input := ""

	for i := 0; i+3 < len(input); i++ {
		m := make(map[byte]bool)
		m[input[i]] = true
		m[input[i+1]] = true
		m[input[i+2]] = true
		m[input[i+3]] = true

		if len(m) == 4 {
			fmt.Println(i + 4)
			return
		}
	}
}

func part2() {
	input := ""

	for i := 0; i+13 < len(input); i++ {
		m := make(map[byte]bool)
		end := i + 13
		for j := i; j < len(input) && j <= end; j++ {
			m[input[j]] = true
		}

		fmt.Println(i, string(input[i]), len(m))
		if len(m) == 14 {
			fmt.Println(i + 14)
			return
		}
	}
}

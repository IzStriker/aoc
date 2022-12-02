package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	max := -1
	current := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		// fmt.Println(value)
		if err != nil {
			if current > max {
				max = current
			}
			current = 0
		} else {
			current += value
		}
	}
	fmt.Println(max)
}

func part2() {
	nums := []int{-1, -1, -1}
	current := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			for i := 0; i < len(nums); i++ {
				if current > nums[i] {
					nums = append(nums[:i+1], nums[i:]...)
					nums[i] = current
					nums = nums[:3]
					break
				}
			}
			current = 0
		} else {
			current += value
		}
	}
	fmt.Println(nums[0] + nums[1] + nums[2])
}

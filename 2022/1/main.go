package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

// 209914
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
		// fmt.Println(value)
		if err != nil {
			nums = append(nums, current)
			current = 0
		} else {
			current += value
		}
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println(nums[0] + nums[1] + nums[2])
}

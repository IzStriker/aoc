package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		part1 := text[:len(text)/2]
		part2 := text[len(text)/2:]

		for _, v := range part1 {
			if strings.ContainsRune(part2, v) {
				if v >= 'a' && v <= 'z' {
					sum += int(v - 'a' + 1)
				} else {
					sum += int(v - 'A' + 27)
				}
				break
			}
		}
	}
	fmt.Println(sum)
}

func part2() {
	file, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	sum := 0
	for i := 0; i+2 < len(lines); i += 3 {
		for _, v := range lines[i] {

			if strings.ContainsRune(lines[i+1], v) && strings.ContainsRune(lines[i+2], v) {
				if v >= 'a' && v <= 'z' {
					sum += int(v - 'a' + 1)
					fmt.Println(string(v), sum)
				} else {
					sum += int(v - 'A' + 27)
					fmt.Println(string(v), sum)

				}
				break
			}
		}
	}
	fmt.Println(sum)
}

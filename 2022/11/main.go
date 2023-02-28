package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type monkey struct {
	number     int
	startItems []int
	operation  []string
	test       int
	throwTrue  int
	throwFalse int
}

func main() {
	m := read()
	part1(m)
}

func read() map[int]monkey {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	monkeys := make(map[int]monkey, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		m := monkey{}
		re := regexp.MustCompile(`\d+`)
		// number
		value, err := strconv.Atoi(re.FindAllString(scanner.Text(), -1)[0])
		if err != nil {
			panic(err)
		}
		m.number = value

		// starting items
		scanner.Scan()
		items := strings.ReplaceAll(scanner.Text(), "Starting items:", "")
		for _, value := range strings.Split(items, ",") {

			value, err := strconv.Atoi(strings.TrimSpace(value))
			if err != nil {
				panic(err)
			}
			m.startItems = append(m.startItems, value)
		}

		// Operations
		scanner.Scan()
		ops := strings.ReplaceAll(scanner.Text(), "  Operation: new = ", "")
		m.operation = strings.Split(ops, " ")

		// test
		scanner.Scan()
		value, err = strconv.Atoi(re.FindAllString(scanner.Text(), -1)[0])
		if err != nil {
			panic(err)
		}
		m.test = value

		// test true
		scanner.Scan()
		value, err = strconv.Atoi(re.FindAllString(scanner.Text(), -1)[0])
		if err != nil {
			panic(err)
		}
		m.throwTrue = value

		// test false
		scanner.Scan()
		value, err = strconv.Atoi(re.FindAllString(scanner.Text(), -1)[0])
		if err != nil {
			panic(err)
		}
		m.throwFalse = value
		monkeys[m.number] = m
		scanner.Scan()
	}
	return monkeys
}

func part1(monkeys map[int]monkey) {
	for i := 0; i < len(monkeys); i++ {
		for _, v := range monkeys[i].startItems {
			monkeys[i]
		}
	}
}

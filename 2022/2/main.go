package main

import (
	"bufio"
	"fmt"
	"os"
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

	total := 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		player1 := int(input[0][0] - 'A' + 1)
		player2 := int(input[1][0] - 'X' + 1)
		score := player2
		win := 0

		if player1 == player2 {
			win = 0
			score += 3
		} else if player2 == 1 && player1 == 3 ||
			player2 == 3 && player1 == 2 ||
			player2 == 2 && player1 == 1 {
			win = 1
			score += 6
		} else {
			win = -1
		}

		total += score
		fmt.Println(player1, player2, win, score, total)
	}
	fmt.Println(total)
}

func part2() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		player1 := int(input[0][0] - 'A' + 1)
		win := int(input[1][0]-'X') - 1
		player2 := -1

		if win > 0 {
			if player1 == 1 {
				player2 = 2
			} else if player1 == 2 {
				player2 = 3
			} else if player1 == 3 {
				player2 = 1
			}
		} else if win == 0 {
			player2 = player1
		} else if win < 0 {
			if player1 == 1 {
				player2 = 3
			} else if player1 == 2 {
				player2 = 1
			} else if player1 == 3 {
				player2 = 2
			}
		}

		score := player2
		if player1 == player2 {
			score += 3
		} else if player2 == 1 && player1 == 3 ||
			player2 == 3 && player1 == 2 ||
			player2 == 2 && player1 == 1 {
			score += 6
		}

		total += score
		fmt.Println(player1, player2, win, score, total)
		// fmt.Println(player1, player2, win)
	}
	fmt.Println(total)
}

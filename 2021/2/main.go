package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    part2()
}

func part1() {
    file, err := os.Open("input.txt")
    
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var h int = 0
    var d int = 0
    for scanner.Scan() {
        temp := scanner.Text()

        instruct := strings.Fields(temp)
        
        switch instruct[0] {
        case "forward":
            val, _ := strconv.Atoi(instruct[1])
            h += val
        case "up":
            val, _ := strconv.Atoi(instruct[1])
            d -= val
        case "down":
            val, _ := strconv.Atoi(instruct[1])
            d += val
        }
    }

    fmt.Println("part 1:", h * d)
}

func part2() {
    file, err := os.Open("input.txt")
    
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var h int = 0
    var d int = 0
    var a int = 0
    for scanner.Scan() {
        temp := scanner.Text()

        instruct := strings.Fields(temp)
        
        switch instruct[0] {
        case "forward":
            val, _ := strconv.Atoi(instruct[1])
            h += val
            d += a * val
        case "up":
            val, _ := strconv.Atoi(instruct[1])
            a -= val
        case "down":
            val, _ := strconv.Atoi(instruct[1])
            a += val
        }
    }

    fmt.Println("part 2:", h * d)
}
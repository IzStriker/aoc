package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    depth := []int{}
    for scanner.Scan() {
        temp, _ := strconv.Atoi(scanner.Text())
        depth = append(depth, temp)
   }

   increased := 0

    for i := 3; i < len(depth); i++ {
        prev := depth[i - 3] + depth[i - 2] + depth[i - 1]
        curr := depth[i - 2] + depth[i - 1] + depth[i]

        if curr > prev {
            increased++
        }
    }
    fmt.Println(increased)

}
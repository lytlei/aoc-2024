package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		panic("Error opening input file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan() {
		input += scanner.Text() + "\n"
	}

	fmt.Println("Part 1 solution is:", solve1(input))
	fmt.Println("Part 2 solution is:", solve2(input))
}

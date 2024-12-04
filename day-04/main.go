package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("Error opening file!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var input [][]rune

	for scanner.Scan() {
		line := []rune(scanner.Text())
		input = append(input, line)
	}

	fmt.Println("The solution to part 1 is:", solve1(input))
	fmt.Println("The solution to part 2 is:", solve2(input))
}

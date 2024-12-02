package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res1, res2 := 0, 0

	for scanner.Scan() {
		line := scanner.Text()
		reportString := strings.Fields(line)
		var report []int

		for _, val := range reportString {
			valInt, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			report = append(report, valInt)
		}

		if isSafe(report) {
			res1++
		}

		if isSafe2(report) {
			res2++
		}
	}

	fmt.Println("Part 1 Solution:", res1)
	fmt.Println("Part 2 Solution:", res2)
}

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
		panic("Error opening the input file!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var prereqs [][2]int
	var updates [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		prereq := [2]int{0, 0}
		prereq[0], err = strconv.Atoi(strings.Split(line, "|")[0])
		if err != nil {
			panic("Error splitting string")
		}
		prereq[1], err = strconv.Atoi(strings.Split(line, "|")[1])
		if err != nil {
			panic("Error splitting string")
		}
		prereqs = append(prereqs, prereq)
	}

	for scanner.Scan() {
		line := scanner.Text()
		updates_str := strings.Split(line, ",")
		var updates_int []int

		for _, val := range updates_str {
			val_int, err := strconv.Atoi(val)
			if err != nil {
				panic("Error converting string to int")
			}
			updates_int = append(updates_int, val_int)
		}

		updates = append(updates, updates_int)
	}

	partOneRes, notOrdered := solve1(prereqs, updates)
	fmt.Println("The solution to part 1 is:", partOneRes)
	fmt.Println("The solution to part 2 is:", solve2(prereqs, notOrdered))

}

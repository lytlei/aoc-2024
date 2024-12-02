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
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	var nums1 []int
	var nums2 []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}

		nums1 = append(nums1, num1)
		nums2 = append(nums2, num2)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Part 1 Solution:", solve1(nums1, nums2))
	fmt.Println("Part 2 Solution:", solve2(nums1, nums2))
}

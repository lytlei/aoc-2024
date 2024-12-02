package main

import (
	"math"
	"slices"
)

func solve1(nums1 []int, nums2 []int) int {
	res := 0

	slices.Sort(nums1)
	slices.Sort(nums2)

	for i := range nums1 {
		res += int(math.Abs(float64(nums1[i] - nums2[i])))
	}

	return res
}

func solve2(nums1 []int, nums2 []int) int {
	res := 0
	m := make(map[int]int)

	for _, num := range nums2 {
		m[num]++
	}

	for _, num := range nums1 {
		res += num * m[num]
	}

	return res
}

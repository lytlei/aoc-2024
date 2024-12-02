package main

import (
	"math"
)

func isSafe(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	isIncreasing := nums[1] > nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return false
		}
		if nums[i] > nums[i-1] && !isIncreasing {
			return false
		}
		if nums[i] < nums[i-1] && isIncreasing {
			return false
		}
		diff := int(math.Abs(float64(nums[i] - nums[i-1])))

		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

// tolerates a single bad level
func isSafe2(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	removedIndex := -1

	// check if it is possible to make it increasing with one removal
	for i := 1; i < len(nums); i++ {
		var prevVal int

		if i-1 == removedIndex {
			prevVal = nums[i-2]
		} else {
			prevVal = nums[i-1]
		}

		diff := int(math.Abs(float64(nums[i] - prevVal)))
		if nums[i] <= prevVal || diff < 1 || diff > 3 { // remove greedily
			if removedIndex != -1 {
				removedIndex = -2
				break
			}

			if i == 1 {
				diff2 := int(math.Abs(float64(nums[0] - nums[2])))

				if nums[2] >= nums[0] && diff2 >= 1 && diff2 <= 3 {
					removedIndex = 1
				} else {
					removedIndex = 0
				}
				continue
			}

			diff2 := int(math.Abs(float64(nums[i] - nums[i-2])))

			if nums[i] > nums[i-2] && diff2 >= 1 && diff2 <= 3 {
				removedIndex = i - 1
			} else {
				removedIndex = i
			}
		}
	}

	if removedIndex != -2 {
		return true
	}

	removedIndex = -1

	// check if it is possible to make it decreasing with one removal
	for i := 1; i < len(nums); i++ {
		var prevVal int

		if i-1 == removedIndex {
			prevVal = nums[i-2]
		} else {
			prevVal = nums[i-1]
		}

		diff := int(math.Abs(float64(nums[i] - prevVal)))
		if nums[i] >= prevVal || diff < 1 || diff > 3 { // remove greedily
			if removedIndex != -1 {
				removedIndex = -2
				break
			}

			if i == 1 {
				diff2 := int(math.Abs(float64(nums[0] - nums[2])))

				if nums[2] <= nums[0] && diff2 >= 1 && diff2 <= 3 {
					removedIndex = 1
				} else {
					removedIndex = 0
				}
				continue
			}

			diff2 := int(math.Abs(float64(nums[i] - nums[i-2])))

			if nums[i] < nums[i-2] && diff2 >= 1 && diff2 <= 3 {
				removedIndex = i - 1
			} else {
				removedIndex = i
			}
		}
	}

	if removedIndex != -2 {
		return true
	}

	return false
}

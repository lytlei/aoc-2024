package main

import (
	"unicode"
)

func solve1(s string) int {
	res := 0
	match := "mul(X,Y)"
	pos := 0
	num1 := 0
	num2 := 0

	for _, c := range s {
		if pos > 0 && match[pos-1] == 'X' && unicode.IsDigit(c) {
			num1 *= 10
			num1 += int(c - rune('0'))
			continue
		} else if pos > 0 && match[pos-1] == 'Y' && unicode.IsDigit(c) {
			num2 *= 10
			num2 += int(c - rune('0'))
			continue
		}

		if match[pos] == 'X' && unicode.IsDigit(c) {
			num1 *= 10
			num1 += int(c - rune('0'))
			pos++
		} else if match[pos] == 'Y' && unicode.IsDigit(c) {
			num2 *= 10
			num2 += int(c - rune('0'))
			pos++
		} else if rune(match[pos]) == c {
			pos++
		} else {
			num1 = 0
			num2 = 0
			pos = 0
		}

		// found matching operation
		if pos == len(match) {
			res += num1 * num2
			pos = 0
			num1 = 0
			num2 = 0
		}
	}

	return res
}

func solve2(s string) int {
	res := 0
	mul := "mul(X,Y)"
	do := "do()"
	dont := "don't()"
	mul_pos := 0
	do_pos := 0
	dont_pos := 0
	active := true
	num1 := 0
	num2 := 0

	for _, c := range s {
		// handle do matching
		if rune(do[do_pos]) == c {
			do_pos++
		} else {
			do_pos = 0
		}

		// handle don't matching
		if rune(dont[dont_pos]) == c {
			dont_pos++
		} else {
			dont_pos = 0
		}

		// handle mul matching
		if mul_pos > 0 && mul[mul_pos-1] == 'X' && unicode.IsDigit(c) {
			num1 *= 10
			num1 += int(c - rune('0'))
			continue
		} else if mul_pos > 0 && mul[mul_pos-1] == 'Y' && unicode.IsDigit(c) {
			num2 *= 10
			num2 += int(c - rune('0'))
			continue
		}

		if mul[mul_pos] == 'X' && unicode.IsDigit(c) {
			num1 *= 10
			num1 += int(c - rune('0'))
			mul_pos++
		} else if mul[mul_pos] == 'Y' && unicode.IsDigit(c) {
			num2 *= 10
			num2 += int(c - rune('0'))
			mul_pos++
		} else if rune(mul[mul_pos]) == c {
			mul_pos++
		} else {
			num1 = 0
			num2 = 0
			mul_pos = 0
		}

		// handle a fully matched string
		if do_pos == len(do) {
			active = true
			do_pos = 0
		}

		if dont_pos == len(dont) {
			active = false
			dont_pos = 0
		}

		if active && mul_pos == len(mul) {
			res += num1 * num2
			mul_pos = 0
			num1 = 0
			num2 = 0
		} else if mul_pos == len(mul) {
			mul_pos = 0
			num1 = 0
			num2 = 0
		}
	}

	return res
}

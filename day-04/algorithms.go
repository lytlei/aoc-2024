package main

var moves = [8][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}
var match = [4]rune{'X', 'M', 'A', 'S'}

func solve1(letters [][]rune) int {
	res := 0

	m := len(letters)
	n := len(letters[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			for _, move := range moves {
				cur_i, cur_j := i, j
				pos := 0
				for {
					if pos < len(match) && cur_i >= 0 && cur_i < m && cur_j >= 0 && cur_j < n && match[pos] == letters[cur_i][cur_j] {
						pos++
					} else {
						break
					}
					cur_i += move[0]
					cur_j += move[1]
				}
				if pos == len(match) {
					res++
				}
			}

		}
	}

	return res
}

func solve2(letters [][]rune) int {
	m, n, res := len(letters), len(letters[0]), 0

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if letters[i][j] == 'A' {
				leftCross := (letters[i-1][j-1] == 'M' && letters[i+1][j+1] == 'S') || (letters[i-1][j-1] == 'S' && letters[i+1][j+1] == 'M')
				rightCross := (letters[i+1][j-1] == 'M' && letters[i-1][j+1] == 'S') || (letters[i+1][j-1] == 'S' && letters[i-1][j+1] == 'M')

				if leftCross && rightCross {
					res++
				}
			}
		}
	}
	return res
}

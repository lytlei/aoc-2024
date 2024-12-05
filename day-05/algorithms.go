package main

func solve1(prereqs [][2]int, updates [][]int) (int, [][]int) {
	m, res := make(map[int][]int), 0
	var notOrdered [][]int

	for _, prereq := range prereqs {
		if m[prereq[1]] == nil {
			m[prereq[1]] = []int{}
		}
		m[prereq[1]] = append(m[prereq[1]], prereq[0])
	}

	for _, update := range updates {
		illegal := make(map[int]bool)
		valid := true
		for _, num := range update {
			_, ok := illegal[num]
			if ok {
				notOrdered = append(notOrdered, update)
				valid = false
				break
			} else {
				for _, notAllowed := range m[num] {
					illegal[notAllowed] = true
				}
			}
		}

		if valid {
			res += update[len(update)/2]
		}
	}

	return res, notOrdered
}

func solve2(prereqs [][2]int, updates [][]int) int {
	res := 0

	for _, update := range updates {
		updateSet := make(map[int]bool)
		for _, val := range update {
			updateSet[val] = true
		}
		m := make(map[int][]int)
		for _, prereq := range prereqs {
			if !updateSet[prereq[0]] || !updateSet[prereq[1]] {
				continue
			}
			if m[prereq[0]] == nil {
				m[prereq[0]] = []int{}
			}
			m[prereq[0]] = append(m[prereq[0]], prereq[1])
		}
		var nums []int
		var correctUpdate []int
		for _, val := range update {
			nums = append(nums, val)
		}
		g := &Graph{
			edges:    m,
			vertices: nums,
		}

		sortedUpdate := g.topologicalSort()

		for _, val := range sortedUpdate {
			if updateSet[val] {
				correctUpdate = append(correctUpdate, val)
			}
		}

		res += correctUpdate[len(correctUpdate)/2]
	}
	return res
}

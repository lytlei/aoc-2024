// tsort algorithm taken from https://reintech.io/blog/topological-sorting-in-go
package main

type Graph struct {
	edges    map[int][]int
	vertices []int
}

func (g *Graph) addEdge(u, v int) {
	g.edges[u] = append(g.edges[u], v)
}

func (g *Graph) topologicalSortUtil(v int, visited map[int]bool, stack *[]int) {
	visited[v] = true

	for _, u := range g.edges[v] {
		if !visited[u] {
			g.topologicalSortUtil(u, visited, stack)
		}
	}

	*stack = append([]int{v}, *stack...)
}

func (g *Graph) topologicalSort() []int {
	stack := []int{}
	visited := make(map[int]bool)

	for _, v := range g.vertices {
		if !visited[v] {
			g.topologicalSortUtil(v, visited, &stack)
		}
	}

	return stack
}

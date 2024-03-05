package main

import "fmt"

type Node struct {
	value    int
	visited  bool
	adjacent []*Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%d", n.value)
}

func newNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func makeGraph(n int, computers [][]int) []*Node {
	graph := make([]*Node, n)
	for i := 0; i < n; i++ {
		graph[i] = newNode(i)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if computers[i][j] == 1 {
				graph[i].adjacent = append(graph[i].adjacent, graph[j])
			}
		}
	}

	return graph
}

func dfs(graph []*Node, start *Node) {
	start.visited = true

	for _, node := range start.adjacent {
		if !node.visited {
			dfs(graph, node)
		}
	}
}

func solution(n int, computers [][]int) int {
	graph := makeGraph(n, computers)
	cnt := 0
	for i := 0; i < n; i++ {
		if !graph[i].visited {
			dfs(graph, graph[i])
			cnt++
		}
	}

	return cnt
}

func main() {
	fmt.Println(
		solution(
			3,
			[][]int{
				{1, 1, 0},
				{1, 1, 0},
				{0, 0, 1},
			},
		),
	)
	fmt.Println(
		solution(
			3,
			[][]int{
				{1, 1, 0},
				{1, 1, 1},
				{0, 1, 1},
			},
		),
	)
}

package main

import "fmt"

type Process struct {
	priority int
	index    int
}

func solution(priorities []int, location int) int {

	var queue []Process

	for i, p := range priorities {
		queue = append(queue, Process{
			priority: p,
			index:    i,
		})
	}

	// queue가 비어 있을 때까지

	order := 0
	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		higherPriority := false
		for _, node := range queue {
			if current.priority < node.priority {
				higherPriority = true
			}
		}

		// 우선 순위 높은게 있다면 뒤로가!
		if higherPriority {
			queue = append(queue, current)
		} else { // 없으면 실해시켜
			order++
			if current.index == location { // 실행시켰는데 그게 location과 같은 경우 반환!
				return order
			}
		}

	}

	return order
}

func main() {
	fmt.Println(solution([]int{2, 1, 3, 2}, 2))
	fmt.Println(solution([]int{1, 1, 9, 1, 1, 1}, 0))
}

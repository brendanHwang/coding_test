package main

//type Process struct {
//	priority int
//	index    int
//}
//
//func Solution(priorities []int, location int) int {
//	queue := make([]Process, len(priorities))
//	for i, priority := range priorities {
//		queue[i] = Process{priority, i}
//	}
//
//	var order int
//	for len(queue) > 0 {
//		current := queue[0]
//		queue = queue[1:]
//
//		// 우선순위가 더 높은 프로세스가 있는지 확인
//		foundHigherPriority := false
//		for _, process := range queue {
//			if process.priority > current.priority {
//				foundHigherPriority = true
//				break
//			}
//		}
//
//		if foundHigherPriority {
//			// 현재 프로세스를 큐의 끝에 다시 추가
//			queue = append(queue, current)
//		} else {
//			// 현재 프로세스 실행
//			order++
//			if current.index == location {
//				// 주어진 위치의 프로세스가 실행되면 그 순서 반환
//				return order
//			}
//		}
//	}
//	return order
//}

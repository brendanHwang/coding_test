package main

import (
	"fmt"
	"stack_queue/process"
)

// bridge_passing_truck
//func main() {
//	result := bridge_paasing_truck.Solution(
//		100,
//		100,
//		[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
//	)
//
//	fmt.Println(result)
//}

// process
func main() {
	fmt.Println(process.Solution([]int{2, 1, 3, 2}, 2))       // 출력: 1
	fmt.Println(process.Solution([]int{1, 1, 9, 1, 1, 1}, 0)) // 출력: 5
}

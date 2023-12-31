package main

import (
	"fmt"
	"sort"
)

type data struct {
	emergency int
	index	  int
}

func solution(emergency []int) []int {
	dataArr := make([]data, len(emergency))

	for i := range emergency {
		dataArr[i] = data{emergency[i], i}
	}

	sort.Slice(dataArr, func(i, j int) bool {
		return dataArr[i].emergency > dataArr[j].emergency 
	}) // 내림차순
	
	answer := make([]int, len(emergency))

	for i := range dataArr {
		answer[dataArr[i].index] = i + 1
	}
	
	return answer
}

func main() {
	emergency := []int{3, 76, 24}
	fmt.Println(solution(emergency))
}

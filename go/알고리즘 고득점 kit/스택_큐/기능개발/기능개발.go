package 기능개발

import "fmt"

func increment(progresses []int, speeds []int) {
	for i, _ := range progresses {
		progresses[i] += speeds[i]
	}
}

func checker(progresses []int, topIdx int) bool {
	return progresses[topIdx] >= 100
}

func solution(progresses []int, speeds []int) []int {

	topIdx := 0
	var completeCnt int
	var ans []int
	for {
		completeCnt = 0
		increment(progresses, speeds)
		if checker(progresses, topIdx) {
			for checker(progresses, topIdx) {
				completeCnt++
				topIdx++

				if topIdx >= len(progresses) {
					ans = append(ans, completeCnt)
					return ans
				}
			}
			ans = append(ans, completeCnt)
		}

	}

}

func main() {

	fmt.Println(solution([]int{95, 90, 99, 99, 80, 99}, []int{1, 1, 1, 1, 1, 1}))

}

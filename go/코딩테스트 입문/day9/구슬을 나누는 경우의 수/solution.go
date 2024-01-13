package 구슬을_나누는_경우의_수

// solution
func solution(balls int, share int) int {
	answer := 1
	for i := 0; i < share; i++ {
		answer *= balls - i
		answer /= i + 1
	}
	return answer
}

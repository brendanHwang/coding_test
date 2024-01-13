package _차원으로_만들기

func solution(num_list []int, n int) [][]int {
	var result [][]int
	for i := 0; i < len(num_list); i += n {
		result = append(result, num_list[i:i+n])
	}
	return result
}

package 점의_위치_구하기

func Solution(dot []int) int {
	if dot[0] > 0 && dot[1] > 0 {
		return 1
	} else if dot[0] < 0 && dot[1] > 0 {
		return 2
	} else if dot[0] < 0 && dot[1] < 0 {
		return 3
	} else {
		return 4
	}
}

package main

func Solution(n int, lost []int, reserve []int) int {
	students := [32]int{}
	for _, v := range lost {
		students[v] = -1
	}

	for _, v := range reserve {
		students[v] += 1
	}

	for i, v := range students {
		if v == 1 {
			if i-1 >= 0 && students[i-1] == -1 {
				students[i-1] = 0
				students[i] = 0
			} else if i+1 < 32 && students[i+1] == -1 {
				students[i+1] = 0
				students[i] = 0
			}
		}
	}

	notAllowed := 0
	for _, v := range students {
		if v == -1 {
			notAllowed++
		}
	}
	return n - notAllowed
}

package 문자열_뒤집기

func Solution(my_string string, s int, e int) string {
	var s_arr []rune

	for i := e; i >= s; i-- {
		s_arr = append(s_arr, s_arr[i])
	}
	return my_string[:s] + string(s_arr) + my_string[e+1:]

}

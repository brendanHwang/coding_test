package 가위_바위_보

func Solution(rsp string) string {
	winnigMap := map[string]string{
		"2": "0",
		"0": "5",
		"5": "2",
	}

	answer := ""

	for _, v := range rsp {
		answer += winnigMap[string(v)]
	}

	return answer
}

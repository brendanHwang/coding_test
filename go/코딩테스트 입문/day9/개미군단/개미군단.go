package 개미군단

func Solution(hp int) int {
	GENERAL := 5
	SOLDIER := 3
	WARRIOR := 1
	ans := hp / GENERAL
	hp = hp % GENERAL
	ans += hp / SOLDIER
	hp = hp % SOLDIER
	ans += hp / WARRIOR
	hp = hp % WARRIOR
	return ans
}

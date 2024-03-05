package bridge_paasing_truck

// 무게의 합을 알려주는 함수
func totalWeight(bridge []int) int {
	sum := 0
	for _, v := range bridge {
		sum += v
	}
	return sum
}

// 마지막을 제외한한 합을 알려주는 함수
func withoutTailTotalWeight(bridge []int) int {
	sum := 0
	for _, v := range bridge[:len(bridge)-1] {
		sum += v
	}
	return sum
}

// 이동을 담당하는 함수
func move(bridge []int, newTruck int) {
	for i := len(bridge) - 1; i >= 1; i-- {
		bridge[i] = bridge[i-1]
	}
	bridge[0] = newTruck
}

func Solution(bridge_length int, weight int, truck_weights []int) int {

	//대기중인 트럭의 맨 앞 (truck_weights)
	head := 0

	//다리 생성 (다리위에 트럭들)
	bridge := make([]int, bridge_length)

	// 경과한 시간
	time := 0

	for head < len(truck_weights) || totalWeight(bridge) != 0 { // 대기중인 트럭이 있거나 or 다리위에 트럭이 있는 경우

		// 대기중이 트럭과 다리 위에 트럭들의 힙이 (다음에 나갈 트럭을 제외) 제한 무게 이하인지 검사
		var addable bool
		if head >= len(truck_weights) { // 대기중인 트럭이 없다면 추가 할게 없으므르 false
			addable = false
		} else {
			addable = withoutTailTotalWeight(bridge)+truck_weights[head] <= weight
		}

		if addable {
			move(bridge, truck_weights[head])
			head++
		} else {
			move(bridge, 0)
		}
		time++
	}

	return time
}

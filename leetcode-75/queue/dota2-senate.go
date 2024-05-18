package queue

func predictPartyVictory(senate string) string {
	radiants := make([]int, 0)
	dires := make([]int, 0)
	for i := 0; i < len(senate); i++ {
		if senate[i] == 'R' {
			radiants = append(radiants, i)
		} else {
			dires = append(dires, i)
		}
	}
	for len(radiants) > 0 && len(dires) > 0 {
		if radiants[0] < dires[0] {
			radiants = append(radiants, radiants[0]+len(senate))
		} else {
			dires = append(dires, dires[0]+len(senate))
		}
		radiants = radiants[1:]
		dires = dires[1:]
	}
	if len(radiants) > 0 {
		return "Radiant"
	}
	return "Dire"
}

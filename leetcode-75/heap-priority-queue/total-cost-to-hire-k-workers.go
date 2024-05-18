package heap_priority_queue

import "math"

func totalCost(costs []int, k int, candidates int) int64 {
	vis := make([]int, len(costs))
	fCosts := make([]int, 0)
	fIndexCosts := make([]int, 0)
	rCosts := make([]int, 0)
	rIndexCosts := make([]int, 0)
	var res int64
	for k > 0 {
		for i := 0; i < len(costs); i++ {
			if len(fCosts) == candidates {
				break
			}
			if vis[i] == 0 {
				fCosts = append(fCosts, costs[i])
				fIndexCosts = append(fIndexCosts, i)
			}
		}

		for i := len(costs) - 1; i >= 0; i-- {
			if len(rCosts) == candidates {
				break
			}
			if vis[i] == 0 {
				rCosts = append(rCosts, costs[i])
				rIndexCosts = append(rIndexCosts, i)
			}
		}

		index, minCost := getMinCost(fCosts, fIndexCosts, rCosts, rIndexCosts)
		for i := 0; i < len(fIndexCosts); i++ {
			if fIndexCosts[i] == index {
				fIndexCosts = append(fIndexCosts[:i], fIndexCosts[i+1:]...)
				fCosts = append(fCosts[:index], fCosts[index+1:]...)
			}
		}
		for i := 0; i < len(rIndexCosts); i++ {
			if rIndexCosts[i] == index {
				rIndexCosts = append(rIndexCosts[:i], rIndexCosts[i+1:]...)
				rCosts = append(rCosts[:index], rCosts[index+1:]...)
			}
		}
		vis[index] = 1
		res += int64(minCost)
		k--
	}
	return res
}

func getMinCost(fCosts, fIndexCosts, rCosts, rIndexCosts []int) (int, int) {
	var index int
	minCost := math.MaxInt
	for i := 0; i < len(fCosts); i++ {
		if minCost > fCosts[i] {
			minCost = fCosts[i]
			index = fIndexCosts[i]
		}
	}
	for i := len(rCosts) - 1; i >= 0; i-- {
		if minCost > rCosts[i] {
			minCost = rCosts[i]
			index = rIndexCosts[i]
		}
	}
	return index, minCost
}

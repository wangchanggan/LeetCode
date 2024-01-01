package backtrack

func combinationSum(candidates []int, target int) [][]int {
	if candidates == nil || len(candidates) == 0 {
		return nil
	}

	sumArr := make([]int, 0)
	res := make([][]int, 0)
	for i := 0; i < len(candidates); i++ {
		combinationSumFindTarget(candidates, sumArr, i, 0, target, &res)
	}
	return res
}

func combinationSumFindTarget(candidates, sumArr []int, index, sum, target int, res *[][]int) {
	if sum > target {
		return
	}
	sum += candidates[index]
	sumArr = append(sumArr, candidates[index])
	if sum == target {
		*res = append(*res, sumArr)
		return
	} else if sum > target {
		return
	} else {
		for i := index; i < len(candidates); i++ {
			combinationSumFindTarget(candidates, deepCopySlice(sumArr), i, sum, target, res)
		}
	}
}

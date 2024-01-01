package backtrack

func permute(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	res := make([][]int, 0)
	vis := make([]int, len(nums))
	n := make([]int, 0)
	for i := 0; i < len(vis); i++ {
		permuteNums(nums, deepCopySlice(vis), deepCopySlice(n), i, &res)
	}
	return res
}

func permuteNums(nums, vis, n []int, index int, res *[][]int) {
	if vis[index] == 0 {
		n = append(n, nums[index])
		vis[index] = 1
		if len(n) == len(nums) {
			*res = append(*res, n)
			return
		}
		for i := 0; i < len(vis); i++ {
			if i != index {
				permuteNums(nums, deepCopySlice(vis), deepCopySlice(n), i, res)
			}
		}
	}
}

func deepCopySlice(slice []int) []int {
	res := make([]int, len(slice))
	for i := 0; i < len(slice); i++ {
		res[i] = slice[i]
	}
	return res
}

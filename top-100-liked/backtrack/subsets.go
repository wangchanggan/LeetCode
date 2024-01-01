package backtrack

func subsets(nums []int) [][]int {
	if nums == nil {
		return nil
	}
	res := make([][]int, 0)
	res = append(res, []int{})
	vis := make([]int, len(nums))
	n := make([]int, 0)
	for i := 0; i < len(vis); i++ {
		subset(nums, vis, n, i, &res)
	}
	return res
}

func subset(nums, vis, n []int, index int, res *[][]int) {
	if vis[index] == 0 {
		n = append(n, nums[index])
		*res = append(*res, n)
		vis[index] = 1
		for i := index + 1; i < len(vis); i++ {
			subset(nums, deepCopySlice(vis), deepCopySlice(n), i, res)
		}
	}
}

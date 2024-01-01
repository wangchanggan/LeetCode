package skill

import "sort"

func nextPermutation(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	} else if len(nums) < 3 {
		sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	} else {
		numsPermutation(nums, 0)
	}
}

func numsPermutation(nums []int, originIndex int) {
	if originIndex == len(nums)-1 {
		sort.Ints(nums)
	} else {
		indexMax := getIndexMax(nums, originIndex)
		if indexMax != originIndex {
			max := nums[indexMax]
			for i := indexMax; i > originIndex+1; i-- {
				nums[i] = nums[i-1]
			}
			nums[originIndex+1] = max
		} else {
			originIndex++
			numsPermutation(nums, originIndex)
		}
	}
}

func getIndexMax(nums []int, originIndex int) int {
	max := nums[originIndex]
	indexMax := originIndex
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			indexMax = i
		}
	}
	return indexMax
}

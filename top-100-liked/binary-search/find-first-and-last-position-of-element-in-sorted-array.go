package binary_search

func searchRange(nums []int, target int) []int {
	if nums == nil {
		return nums
	} else if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	res := make([]int, 2)
	res[0] = -1
	res[1] = -1
	findFirstAndLastPositionOfElementInSortedArray(nums, 0, len(nums)-1, target, res)
	return res
}

func findFirstAndLastPositionOfElementInSortedArray(nums []int, front, rear, target int, res []int) {
	if front < 0 || front >= len(nums) || rear < 0 || rear >= len(nums) || front > rear {
		return
	}
	if nums[front] == target && nums[rear] == target {
		res[0] = front
		res[1] = rear
		return
	} else if nums[front] == target && nums[rear] != target {
		res[0] = front
		findFirstAndLastPositionOfElementInSortedArray(nums, front, rear-1, target, res)
	} else if nums[front] != target && nums[rear] == target {
		res[1] = rear
		findFirstAndLastPositionOfElementInSortedArray(nums, front+1, rear, target, res)
	} else {
		findFirstAndLastPositionOfElementInSortedArray(nums, front+1, rear-1, target, res)
	}
}

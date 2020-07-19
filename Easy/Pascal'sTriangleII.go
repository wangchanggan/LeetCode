package Easy

func PascalsTriangleII(rowIndex int) []int {
	var result []int
	nums := make([][]int, rowIndex+1)
	if rowIndex==0{
		result = append(result, 1)
		return result
	}
	nums[0] = append(nums[0], 1)
	if rowIndex==1{
		result = append(result, 1)
		result = append(result, 1)
		return result
	}
	nums[1] = append(nums[1], 1)
	nums[1] = append(nums[1], 1)
	for i:=3;i<=rowIndex+1;i++{
		nums[i-1] = append(nums[i-1],nums[i-2][0])
		for j:=1;j<i-1;j++{
			nums[i-1] = append(nums[i-1],nums[i-2][j-1]+nums[i-2][j])
		}
		nums[i-1] = append(nums[i-1],nums[i-2][i-2])
	}
	for i:=0;i<=rowIndex;i++{
		result = append(result, nums[rowIndex][i])
	}
	return result
}
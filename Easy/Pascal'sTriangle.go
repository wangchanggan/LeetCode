package Easy

func PascalsTriangle(numRows int) [][]int {
	nums := make([][]int, numRows)
	if numRows==0{
		return nums
	}
	nums[0] = append(nums[0], 1)
	if numRows==1{
		return nums
	}
	nums[1] = append(nums[1], 1)
	nums[1] = append(nums[1], 1)
	if numRows==2{
		return nums
	}
	for i:=3;i<=numRows;i++{
		nums[i-1] = append(nums[i-1],nums[i-2][0])
		for j:=1;j<i-1;j++{
			nums[i-1] = append(nums[i-1],nums[i-2][j-1]+nums[i-2][j])
		}
		nums[i-1] = append(nums[i-1],nums[i-2][i-2])
	}
	return nums
}
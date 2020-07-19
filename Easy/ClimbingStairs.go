package Easy

func ClimbingStairs(n int) int {
	if n==0||n==1||n==2||n==3{
		return n
	}
	var nums [10000]int
	nums[2] = 2
	nums[3] = 3
	for i:=4;i<=n;i++{
		nums[i] = nums[i-1]+nums[i-2]
	}
	return nums[n]
}
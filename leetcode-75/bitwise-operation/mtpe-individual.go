package bitwise_operation

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = countB(i)
	}
	return res
}

func countB(n int) int {
	var count int
	for n > 0 {
		if n%2 == 1 {
			count++
		}
		n /= 2
	}
	return count
}

package Hard

import (
	"fmt"
	"math"
	"sort"
)

func MinimumDifference(nums []int) int {
	n := len(nums) / 2
	a := nums[:n]
	res := make([][]int, n+1)
	for i := 0; i < 1<<n; i++ {
		fmt.Println(i)
		sum, cnt := 0, 0
		for j, v := range a {
			if i>>j&1 > 0 { // 1 视作取正
				sum += v
				cnt++
			} else { // 0 视作取负
				sum -= v
			}
		}
		res[cnt] = append(res[cnt], sum) // 按照取正的个数将元素和分组
	}
	fmt.Println(res)

	for _, b := range res {
		sort.Ints(b) // 排序，方便下面二分
	}
	fmt.Println(res)

	ans := math.MaxInt64
	a = nums[n:]
	for i := 0; i < 1<<n; i++ {
		sum, cnt := 0, 0
		for j, v := range a {
			if i>>j&1 > 0 {
				sum += v
				cnt++
			} else {
				sum -= v
			}
		}
		// 在对应的组里二分最近的数
		b := res[cnt]
		j := sort.SearchInts(b, sum)
		if j < len(b) {
			ans = min(ans, b[j]-sum)
		}
		if j > 0 {
			ans = min(ans, sum-b[j-1])
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 9))
}

func combinationSum3(k int, n int) [][]int {
	if n <= k {
		return nil
	}
	res := new([][]int)
	list := make([]int, 0)
	dfsCombinationSum3(k, 0, 1, n, list, res)
	return *res
}

func dfsCombinationSum3(k, c, m, n int, list []int, res *[][]int) {
	if k == 0 {
		if c == n {
			*res = append(*res, list)
		}
		return
	}
	if k < 0 || c > n {
		return
	}
	for i := m; i <= 9; i++ {
		if c+i <= n {
			newList := newSlice(list)
			newList = append(newList, i)
			dfsCombinationSum3(k-1, c+i, i+1, n, newList, res)
		} else {
			break
		}
	}
}

func newSlice(a []int) []int {
	b := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		b[i] = a[i]
	}
	return b
}

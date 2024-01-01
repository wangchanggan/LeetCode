package heap

import "sort"

type MedianFinder struct {
	Nums []int
}

func Constructor() MedianFinder {
	var mf MedianFinder
	mf.Nums = make([]int, 0)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	this.Nums = append(this.Nums, num)
	sort.Ints(this.Nums)
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.Nums)%2 == 0 {
		index := len(this.Nums) / 2
		return float64(this.Nums[index-1]+this.Nums[index]) / 2
	} else {
		return float64(this.Nums[len(this.Nums)/2])
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

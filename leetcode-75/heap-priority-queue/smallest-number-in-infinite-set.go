package heap_priority_queue

import "sort"

type SmallestInfiniteSet struct {
	min  int
	nums []int
}

func Constructor() SmallestInfiniteSet {
	var smallestInfiniteSet SmallestInfiniteSet
	smallestInfiniteSet.min = 1
	smallestInfiniteSet.nums = make([]int, 0)
	return smallestInfiniteSet
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	if len(this.nums) == 0 {
		smallestNum := this.min
		this.min++
		return smallestNum
	}

	smallestNum := this.nums[0]
	this.nums = this.nums[1:]
	return smallestNum
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	for _, n := range this.nums {
		if n == num {
			return
		}
	}
	if num < this.min {
		this.nums = append(this.nums, num)
		sort.Ints(this.nums)
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

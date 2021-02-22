/*Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Implement KthLargest class:

KthLargest(int k, int[] nums) Initializes the object with the integer k and the stream of integers nums.
int add(int val) Returns the element representing the kth largest element in the stream.


Example 1:

Input
["KthLargest", "add", "add", "add", "add", "add"]
[[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
Output
[null, 4, 5, 5, 8, 8]

Explanation
KthLargest kthLargest = new KthLargest(3, [4, 5, 8, 2]);
kthLargest.add(3);   // return 4
kthLargest.add(5);   // return 5
kthLargest.add(10);  // return 5
kthLargest.add(9);   // return 8
kthLargest.add(4);   // return 8


Constraints:

1 <= k <= 104
0 <= nums.length <= 104
-104 <= nums[i] <= 104
-104 <= val <= 104
At most 104 calls will be made to add.
It is guaranteed that there will be at least k elements in the array when you search for the kth element.*/

package Easy

import "sort"

type KthLargest struct {
	K          int
	SortedNums []int
}

func NewKthLargest(k int, nums []int) KthLargest {
	var kthLargest KthLargest
	kthLargest.K = k
	sort.Ints(nums)
	kthLargest.SortedNums = nums
	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	lenght := len(this.SortedNums)
	var flag bool
	for i := 0; i < lenght; i++ {
		if val <= this.SortedNums[i] {
			newSortedNums := make([]int, 0)
			for j := 0; j < i; j++ {
				newSortedNums = append(newSortedNums, this.SortedNums[j])
			}
			newSortedNums = append(newSortedNums, val)
			for j := i; j < lenght; j++ {
				newSortedNums = append(newSortedNums, this.SortedNums[j])
			}
			flag = true
			this.SortedNums = newSortedNums
			break
		}
	}
	if !flag {
		this.SortedNums = append(this.SortedNums, val)
	}

	return this.SortedNums[len(this.SortedNums)-this.K]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

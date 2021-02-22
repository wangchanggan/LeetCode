package Easy

import (
	"fmt"
	"testing"
)

func TestKthLargestElementinaStream(t *testing.T) {
	kthLargest := NewKthLargest(3, []int{4, 5, 8, 2})
	fmt.Println(kthLargest)
	fmt.Println(kthLargest.Add(3))
	fmt.Println(kthLargest)
	fmt.Println(kthLargest.Add(5))
	fmt.Println(kthLargest)
	fmt.Println(kthLargest.Add(10))
	fmt.Println(kthLargest.Add(9))
	fmt.Println(kthLargest.Add(4))
}

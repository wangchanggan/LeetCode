package Easy

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int)  {
	var count int
	for i:=m;i<len(nums1);i++{
		nums1[i] = nums1[count]
		count++
	}
	for i:=0;i<len(nums1)-1;i++{
		for j:=i+1;j<len(nums1);j++{
			if nums1[i]>nums1[j]{
				tmp := nums1[j]
				nums1[j] = nums1[i]
				nums1[i] = tmp
			}
		}
	}
}
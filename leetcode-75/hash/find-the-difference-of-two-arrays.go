package hash

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1Map := make(map[int]int)
	nums2Map := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]] = 1
	}
	for i := 0; i < len(nums2); i++ {
		nums2Map[nums2[i]] = 1
	}
	res := make([][]int, 2)
	for k, _ := range nums1Map {
		if nums2Map[k] != 1 {
			res[0] = append(res[0], k)
		}
	}
	for k, _ := range nums2Map {
		if nums1Map[k] != 1 {
			res[1] = append(res[1], k)
		}
	}
	return res
}

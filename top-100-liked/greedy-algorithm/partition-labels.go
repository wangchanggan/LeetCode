package greedy_algorithm

func partitionLabels(s string) []int {
	if len(s) == 0 {
		return nil
	}
	lastPosistion := [26]int{}
	for i, c := range s {
		lastPosistion[c-'a'] = i
	}

	start, end := 0, 0
	partition := make([]int, 0)
	for i, c := range s {
		if lastPosistion[c-'a'] > end {
			end = lastPosistion[c-'a']
		}
		if i == end {
			partition = append(partition, end-start+1)
			start = end + 1
		}
	}
	return partition
}

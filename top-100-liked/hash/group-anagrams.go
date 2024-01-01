package hash

func groupAnagrams(strs []string) [][]string {
	if strs == nil {
		return nil
	} else if len(strs) == 1 {
		return [][]string{strs}
	}
	mp := make(map[[26]int][]string)
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

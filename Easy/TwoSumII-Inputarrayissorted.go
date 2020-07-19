package Easy

func twoSum(numbers []int, target int) []int {
	var result []int
	for i:=0;i<len(numbers)-1;i++{
		for j:=i+1;j<len(numbers);j++{
			if numbers[i]+numbers[j]==target{
				result = append(result, i+1)
				result = append(result, j+1)
			}
		}
	}
	return result
}
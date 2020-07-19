package Easy

func max(num1, num2 int) int {
	if num1 > num2{
		return num1
	}
	return num2
}

func BestTimetoBuyandSellStock(prices []int) int {
	if len(prices)==0||len(prices)==1{
		return 0
	}
	maxPrice :=  0
	price := prices[0]
	for i:=1;i<len(prices);i++{
		if maxPrice > prices[i]{
			price = prices[i]
		}else{
			maxPrice = max(prices[i]-price, maxPrice)
		}
	}
	return maxPrice
}
package Easy

func BestTimetoBuyandSellStockII(prices []int) int {
	var result int
	for i := 1; i < len(prices); i++ {
		if(prices[i] > prices[i-1]) {
			result += prices[i] - prices[i-1]
		}
	}
	return result
}
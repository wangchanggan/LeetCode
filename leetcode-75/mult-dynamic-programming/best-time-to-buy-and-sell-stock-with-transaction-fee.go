package mult_dynamic_programming

func maxProfit(prices []int, fee int) int {
	if prices == nil || len(prices) == 0 {
		return 0
	}
	sell := 0
	buy := -prices[0]
	for i := 1; i < len(prices); i++ {
		sell = max(sell, sell+prices[i]-fee)
		sell = max(buy, buy-prices[i])
	}
	return sell
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

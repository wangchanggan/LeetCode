package stack

type StockSpanner struct {
	prices []int
}

func Constructor() StockSpanner {
	var stockSpanner StockSpanner
	stockSpanner.prices = make([]int, 0)
	return stockSpanner
}

func (this *StockSpanner) Next(price int) int {
	var count int
	for i := len(this.prices) - 1; i >= 0; i-- {
		if price < this.prices[i] {
			count = len(this.prices) - i
			break
		}
	}
	this.prices = append(this.prices, price)
	if count == 0 {
		return len(this.prices)
	}
	return count
}

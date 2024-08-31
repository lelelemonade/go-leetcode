package problem901

type StockSpanner struct {
	prices []int
}

func Constructor() StockSpanner {
	return StockSpanner{prices: make([]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	spanStart := -1
	for i := len(this.prices) - 1; i >= 0; i-- {
		if this.prices[i] > price {
			spanStart = i
			break
		}
	}
	this.prices = append(this.prices, price)
	if len(this.prices) == -1 {
		return len(this.prices)
	}
	return len(this.prices) - spanStart - 1
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

package buysell

func MaxProfit(prices []int) int {
	buy := 0
	sell := 1
	profit := 0
	for sell < len(prices) {
		currProfit := prices[sell] - prices[buy]
		if currProfit <= 0 {
			buy, sell = sell, sell+1
		} else if currProfit > profit {
			profit = currProfit
			sell++
		} else {
			sell++
		}

	}
	return profit
}
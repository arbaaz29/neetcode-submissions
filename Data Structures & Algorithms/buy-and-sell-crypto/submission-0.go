func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProf := 0
	for _, price := range prices{
		if price < minPrice{
			minPrice = price
		}
		profit := price - minPrice
		if profit > maxProf {
			maxProf = profit
		}
	}
	return maxProf
}

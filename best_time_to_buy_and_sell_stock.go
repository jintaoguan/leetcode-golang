package main

import "fmt"

func main() {
	stockPrices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(stockPrices))
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxProfit := 0
	minPrice := prices[0]
	for _, price := range prices {
		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}
	return maxProfit
}


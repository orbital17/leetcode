package cooldown

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	var buy, sell, pBuy, pSell, ppSell int
	pBuy = -prices[0]
	for _, p := range prices {
		buy = Max(ppSell-p, pBuy)
		sell = Max(pBuy+p, pSell)
		ppSell, pSell, pBuy = pSell, sell, buy
	}
	return sell
}

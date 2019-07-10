package cooldown

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	var buy, sell, pBuy, pSell, ppSell int
	pBuy = -1
	for _, p := range prices {
		buy = Max(ppSell-p, pBuy)
		sell = Max(pBuy+p, pSell)
		ppSell, pSell, pBuy = pSell, sell, buy
	}
	return sell
}

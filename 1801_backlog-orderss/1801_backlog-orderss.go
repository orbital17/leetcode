package leetcode

import "github.com/emirpasic/gods/maps/treemap"

func getNumberOfBacklogOrders(orders [][]int) int {
	buy := treemap.NewWithIntComparator()
	sell := treemap.NewWithIntComparator()
	for _, order := range orders {
		price, amount, orderType := order[0], order[1], order[2]
		if orderType == 0 { // buy
			it := sell.Iterator()
			it.First()
			for !sell.Empty() {
				k := it.Key().(int)
				v := it.Value().(int)
				if k > price {
					break
				}
				if amount < v {
					sell.Put(k, v-amount)
					amount = 0
					break
				}
				it.Next()
				sell.Remove(k)
				amount -= v
			}
			if amount > 0 {
				v, found := buy.Get(price)
				if !found {
					v = 0
				}
				buy.Put(price, v.(int)+amount)
			}
		} else {
			it := buy.Iterator()
			it.Last()
			for !buy.Empty() {
				k := it.Key().(int)
				v := it.Value().(int)
				if k < price {
					break
				}
				if amount < v {
					buy.Put(k, v-amount)
					amount = 0
					break
				}
				it.Prev()
				buy.Remove(k)
				amount -= v
			}
			if amount > 0 {
				v, found := sell.Get(price)
				if !found {
					v = 0
				}
				sell.Put(price, v.(int)+amount)
			}
		}
	}
	var sum int
	for _, v := range buy.Values() {
		sum += v.(int)
		sum %= 1000000007
	}
	for _, v := range sell.Values() {
		sum += v.(int)
		sum %= 1000000007
	}
	return sum
}

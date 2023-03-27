package price_discount

func CalculateTotalPrice(prices []int, discount int) int {
	majorPrice := prices[0] //to apply it the discount
	sum := 0
	for _, p := range prices {
		if p > majorPrice {
			majorPrice = p
		}
		sum += p
	}
	preSum := sum - majorPrice
	disc := float64(majorPrice) * float64(discount) / 100
	mostExpensiveItemWthDiscount := majorPrice - int(disc)
	total := preSum + mostExpensiveItemWthDiscount
	return total
}

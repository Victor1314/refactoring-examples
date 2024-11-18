
func hasDiscount(order Order) bool {
	basePrice := order.basePrice()
	return basePrice > 1000
}

// def hasDiscount(order):
//     basePrice = order.basePrice()
//     return basePrice > 1000
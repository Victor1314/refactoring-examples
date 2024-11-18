

func hasDiscount(order Order) bool {
	return order.basePrice() > 1000
}

// def hasDiscount(order):
//     return order.basePrice() > 1000
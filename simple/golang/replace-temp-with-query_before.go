



func calculateTotal() {
	basePrice := quantity * itemPrice
	if basePrice > 1000 {
		return basePrice * 0.95
	} else {
		return basePrice * 0.98
	}
}

// def calculateTotal():
//     basePrice = quantity * itemPrice
//     if basePrice > 1000:
//         return basePrice * 0.95
//     else:
//         return basePrice * 0.98
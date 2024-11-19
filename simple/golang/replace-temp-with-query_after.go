
func calculateTotal() {

	if basePrice() > 1000 {
		return basePrice() * 0.95
	} else {
		return basePrice() * 0.98
	}
}

func basePrice() {
	return quantity * itemPrice
}

// def calculateTotal():
//     if basePrice() > 1000:
//         return basePrice() * 0.95
//     else:
//         return basePrice() * 0.98

// def basePrice():
//     return quantity * itemPrice
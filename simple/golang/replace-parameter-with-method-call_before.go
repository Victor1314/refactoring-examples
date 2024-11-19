


// Calling a query method and passing its results as the parameters of another method,
//while that method could call the query directly.

basePrice := quantity * itemPrice
seasonalDiscount := getSeasonalDiscount()
fees := getFees()
finalPrice := discountedPrice(basePrice, seasonalDiscount, fees)

// basePrice = quantity * itemPrice
// seasonalDiscount = self.getSeasonalDiscount()
// fees = self.getFees()
// finalPrice = discountedPrice(basePrice, seasonalDiscount, fees)



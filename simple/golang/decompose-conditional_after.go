


if isSummer(date) {
	charge = summerCharge(quantity)
} else {
	charge = winterCharge(quantity)
}


func isSummer(date time.Time) bool {
	return date.before(SUMMER_START) || date.after(SUMMER_END)
}


func summerCharge(quantity int) int {
	return quantity * summerRate
}   

func winterCharge(quantity int) int {
	return quantity * winterRate + winterServiceCharge
}



// if isSummer(date):
//     charge = summerCharge(quantity)
// else:
//     charge = winterCharge(quantity)
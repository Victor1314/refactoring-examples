


if date.before(SUMMER_START) || date.after(SUMMER_END) {
	charge = quantity*winterRate + winterServiceCharge
} else {
	charge = quantity * summerRate
}

// if date.before(SUMMER_START) or date.after(SUMMER_END):
//     charge = quantity * winterRate + winterServiceCharge
// else:
//     charge = quantity * summerRate
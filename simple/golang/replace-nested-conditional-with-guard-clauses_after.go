


func (e Employee) getPayAmount() int {
	if e.isDead {
		return deadAmount()
	}
	if e.isSeparated {
		return separatedAmount()
	}
	if e.isRetired {
		return retiredAmount()
	}
	return normalPayAmount()
}

// def getPayAmount(self):
//     if self.isDead:
//         return deadAmount()
//     if self.isSeparated:
//         return separatedAmount()
//     if self.isRetired:
//         return retiredAmount()
//     return normalPayAmount()
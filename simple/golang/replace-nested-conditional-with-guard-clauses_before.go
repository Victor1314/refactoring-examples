

func (e Employee) getPayAmount() int {

	var result int
	if e.isDead {
		result = deadAmount()
	} else {
		if e.isSeparated {
			result = separatedAmount()
		} else {
			if e.isRetired {
				result = retiredAmount()
			} else {
				result = normalPayAmount()
			}
		}
	}

	return result
}

// def getPayAmount(self):
//     if self.isDead:
//         result = deadAmount()
//     else:
//         if self.isSeparated:
//             result = separatedAmount()
//         else:
//             if self.isRetired:
//                 result = retiredAmount()
//             else:
//                 result = normalPayAmount()
//     return result

    

    
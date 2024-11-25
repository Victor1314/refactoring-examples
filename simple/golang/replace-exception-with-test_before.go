


func getValueForPeriod(periodNumber int) int {
	// Recover from panic if index is out of range
	defer func() {
		if r := recover(); r != nil {
			result = 0
		}
	}()
	return values[periodNumber]
}

// def getValueForPeriod(periodNumber):
//     try:
//         return values[periodNumber]
//     except IndexError:
//         return 0

func getValueForPeriod(periodNumber int) int {
	if periodNumber >= len(values) {
		return 0
	}
	return values[periodNumber]
}

// def getValueForPeriod(self, periodNumber):
//     if periodNumber >= len(self.values):
//         return 0
//     return self.values[periodNumber]
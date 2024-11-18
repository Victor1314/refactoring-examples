

var (
	seniority      int
	monthsDisabled int
	isPartTime     bool
)

func isNotEligibleForDisability() bool {
	return seniority < 2 || monthsDisabled > 12 || isPartTime
}

func disabilityAmount() int {

	if isNotEligibleForDisability() {
		return 0
	}
	// Compute the disability amount.
	// ...
}

// def disabilityAmount():
//     if isNotEligibleForDisability():
//         return 0
//     # Compute the disability amount.
//     # ...
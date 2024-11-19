



func foundPerson(people []string) {

	candidates := []string{"Don", "John", "Kent"}
	for _, person := range people {
		if slices.Contains(candidates, person) {
			return person
		}
	}
	return ""
}

// def foundPerson(people):
//     candidates = ["Don", "John", "Kent"]
//     return people if people in candidates else ""

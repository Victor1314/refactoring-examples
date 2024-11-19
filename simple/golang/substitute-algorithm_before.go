

func foundPerson(people []string) {

	for i := 0; i < len(people); i++ {
		if people[i] == "Don" {
			return "Don"
		}
		if people[i] == "John" {
			return "John"
		}
		if people[i] == "Kent" {
			return "Kent"
		}
	}
	return ""
}

// def foundPerson(people):
//     for i in range(len(people)):
//         if people[i] == "Don":
//             return "Don"
//         if people[i] == "John":
//             return "John"
//         if people[i] == "Kent":
//             return "Kent"
//     return ""
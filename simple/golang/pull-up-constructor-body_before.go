

// Your subclasses have constructors with code that's mostly identical

type Employee struct {
	Name string
	ID   int
}

type Manager struct {
	Employee
	Grade int
}

func NewManager(name string, id int, grade int) *Manager {
	return &Manager{
		Employee: Employee{
			Name: name,
			ID:   id,
		},
		Grade: grade,
	}
}

// class Manager(Employee):
//     def __init__(self, name, id, grade):
//         self.name = name
//         self.id = id
//         self.grade = grade
//     # ...


type Employee struct {
	Name string
	ID   int
}

func NewEmployee(name string, id int) Employee {
	return Employee{
		Name: name,
		ID:   id,
	}
}

type Manager struct {
	Employee
	Grade int
}

func NewManager(name string, id int, grade int) *Manager {
	return &Manager{
		Employee: NewEmployee(name, id),
		Grade:    grade,
	}
}

// class Manager(Employee):
//     def __init__(self, name, id, grade):
//         Employee.__init__(name, id)
//         self.grade = grade
//     # ...
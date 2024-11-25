package main

import "fmt"


type Customer interface {
	IsNull() bool
	GetPlan() string
}


type NullCustomer struct{}

func (n NullCustomer) IsNull() bool {
	return true
}

func (n NullCustomer) GetPlan() string {
	return "Null Plan"
}


type RegularCustomer struct {
	plan string
}

func (r RegularCustomer) IsNull() bool {
	return false
}

func (r RegularCustomer) GetPlan() string {
	return r.plan
}


type Order struct {
	customer Customer
}

func main() {

	order := Order{customer: nil}


	customer := order.customer
	if customer == nil {
		customer = NullCustomer{}
	}

	plan := customer.GetPlan()
	fmt.Println("Customer Plan:", plan)
}

// class NullCustomer(Customer):

//     def isNull(self):
//         return True

//     def getPlan(self):
//         return self.NullPlan()

//     # Some other NULL functionality.

// # Replace null values with Null-object.
// customer = order.customer or NullCustomer()

// # Use Null-object as if it's normal subclass.
// plan = customer.getPlan()

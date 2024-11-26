

type NullCustomer struct{}


func (n NullCustomer) isNull() bool {
	return true
}

func (n NullCustomer) GetPlan() *BillingPlan {
	return BillingPlanBasic()
}


// # Replace null values with Null-object.

if order.customer == nil {
	customer = NullCustomer{}
} else {
	customer = order.customer
}
// # Use Null-object as if it's normal subclass.
plan = customer.GetPlan()

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

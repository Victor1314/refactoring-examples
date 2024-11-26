


if customer == nil {
	plan = BillingPlanBasic()
} else {
	plan = customer.GetPlan()
}

// if customer is None:
//     plan = BillingPlan.basic()
// else:
//     plan = customer.getPlan()
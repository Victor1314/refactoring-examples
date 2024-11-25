


func (e *Expense) getExpenseLimit() int {
	// Assert that either expenseLimit is set or primaryProject is not nil
	if e.expenseLimit == NULL_EXPENSE && e.primaryProject == nil {
		panic("Either expenseLimit must be set or primaryProject must be non-nil")
	}

	if e.expenseLimit != NULL_EXPENSE {
		return e.expenseLimit
	}
	return e.primaryProject.getMemberExpenseLimit()
}

// def getExpenseLimit(self):
//     assert (self.expenseLimit != NULL_EXPENSE) or (self.primaryProject != None)

//     return self.expenseLimit if (self.expenseLimit != NULL_EXPENSE) else \
//         self.primaryProject.getMemberExpenseLimit()

// For a portion of code to work correctly, certain conditions or values must be true.

func (e *Expense) getExpenseLimit() int {
	// Should have either expense limit or
	// a primary project.
	if e.expenseLimit != NULL_EXPENSE {
		return e.expenseLimit
	}
	return e.primaryProject.getMemberExpenseLimit()
}

// def getExpenseLimit(self):
//     # Should have either expense limit or
//     # a primary project.
//     return self.expenseLimit if self.expenseLimit != NULL_EXPENSE else \
//         self.primaryProject.getMemberExpenseLimit()
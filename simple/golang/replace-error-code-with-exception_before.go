


func (b *BankAccount) withdraw(amount int) int {
	if amount > b.balance {
		return -1
	}
	b.balance -= amount
	return 0
}

// def withdraw(self, amount):
//     if amount > self.balance:
//         return -1
//     else:
//         self.balance -= amount
//     return 0




func (b *BankAccount) withdraw(amount int) error {
	if amount > b.balance {
		return errors.New("insufficient balance")
	}
	b.balance -= amount
	return nil
}

// def withdraw(self, amount):
//     if amount > self.balance:
//         raise BalanceException()
//     self.balance -= amount
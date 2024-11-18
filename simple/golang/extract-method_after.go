
func (order *Order) printOwing() {
	order.printBanner()
	order.printDetails()
}

func (order *Order) printDetails(outstanding float64) {
	fmt.Println("name:", order.name)
	fmt.Println("amount:", order.outstanding)
}

// def printOwing(self):
//     self.printBanner()
//     self.printDetails(self.getOutstanding())

// def printDetails(self, outstanding):
//     print("name:", self.name)
//     print("amount:", outstanding)
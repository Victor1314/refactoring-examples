
func (order *Order) printOwing() {

	order.printBanner()

	// print details
	fmt.Println("name:", order.name)
	fmt.Println("amount:", order.getOutstanding())
}

// def printOwing(self):
//     self.printBanner()

//     # print details
//     print("name:", self.name)
//     print("amount:", self.getOutstanding())



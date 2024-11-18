

// When a method body is more obvious than the method itself, use this technique.

type PizzaDelivery struct {
	numberOfLateDeliveries int
}

func (p *PizzaDelivery) getRating() int {

	if p.moreThanFiveLateDeliveries() {
		return 2
	}
	return 1
}

func (p *PizzaDelivery) moreThanFiveLateDeliveries() bool {
	return p.numberOfLateDeliveries > 5
}

// class PizzaDelivery:
//     # ...
//     def getRating(self):
//         return 2 if self.moreThanFiveLateDeliveries() else 1

//     def moreThanFiveLateDeliveries(self):
//         return self.numberOfLateDeliveries > 5
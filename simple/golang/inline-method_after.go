

type PizzaDelivery struct {
	numberOfLateDeliveries int
}

func (p *PizzaDelivery) getRating() int {
	if p.numberOfLateDeliveries > 5 {
		return 2
	}
	return 1
}

// class PizzaDelivery:
//   # ...
//   def getRating(self):
//     return 2 if self.numberOfLateDeliveries > 5 else 1
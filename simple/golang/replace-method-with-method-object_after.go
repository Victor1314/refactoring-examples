


type PriceCalculator struct {
	primaryBasePrice   int
	secondaryBasePrice int
	tertiaryBasePrice  int
}

// NewPriceCalculator 构造函数
func NewPriceCalculator(order *Order) *PriceCalculator {
	// Copy relevant information from the
	//  order object.
	return &PriceCalculator{
		primaryBasePrice:   0,
		secondaryBasePrice: 0,
		tertiaryBasePrice:  0,
	}
}

func (p *PriceCalculator) Compute() int {
	// Perform long computation.

}

type Order struct {
}

func (o *Order) Price() int {
	return NewPriceCalculator(o).Compute()
}

// class Order:
//     # ...
//     def price(self):
//         return PriceCalculator(self).compute()

// class PriceCalculator:
//     def __init__(self, order):
//         self._primaryBasePrice = 0
//         self._secondaryBasePrice = 0
//         self._tertiaryBasePrice = 0
//         # Copy relevant information from the
//         # order object.

//     def compute(self):
//         # Perform long computation.

// You have a long method in which
// the local variables are so intertwined that you can't apply Extract Method.

type Order struct {
}

func (o *Order) Price() int {
	primaryBasePrice := 0
	secondaryBasePrice := 0
	tertiaryBasePrice := 0
	// Perform long computation.

}

// class Order:
//     # ...
//     def price(self):
//         primaryBasePrice = 0
//         secondaryBasePrice = 0
//         tertiaryBasePrice = 0
//         # Perform long computation.



    
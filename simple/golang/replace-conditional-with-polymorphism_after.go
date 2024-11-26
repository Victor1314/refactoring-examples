


type Bird interface {
	getSpeed() int
}


type European struct {
    // ...
}

func (e European) getSpeed() int {
	return e.getBaseSpeed()
}

type African struct {
    // ...
}

func (a African) getSpeed() int {
	return a.getBaseSpeed() - a.getLoadFactor()*a.numberOfCoconuts
}


// NorwegianBlue implementation
type NorwegianBlue struct {
    // ...
}

func (n NorwegianBlue) getSpeed() int {
	if n.isNailed {
		return 0
	}
	return n.getBaseSpeed()
}

// somewhere in client code
speed := bird.getSpeed()




// class Bird:
//     # ...
//     def getSpeed(self):
//         pass

// class European(Bird):
//     def getSpeed(self):
//         return self.getBaseSpeed()

// class African(Bird):
//     def getSpeed(self):
//         return self.getBaseSpeed() - self.getLoadFactor() * self.numberOfCoconuts

// class NorwegianBlue(Bird):
//     def getSpeed(self):
//         return 0 if self.isNailed else self.getBaseSpeed(self.voltage)

// # Somewhere in client code
// speed = bird.getSpeed()
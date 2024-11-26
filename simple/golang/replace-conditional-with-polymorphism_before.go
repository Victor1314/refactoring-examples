


func (b Bird) getSpeed() int {
	if b.type == "EUROPEAN" {
		return b.getBaseSpeed()
	}else if b.type == "AFRICAN" {
		return b.getBaseSpeed() - b.getLoadFactor() * b.numberOfCoconuts
	}else if b.type == "NORWEGIAN_BLUE" {
		if b.isNailed {
            return 0
        }
        return b.getBaseSpeed(b.voltage)
		
	}else {
		panic("Should be unreachable")
	}   
}






// class Bird:
//     # ...
//     def getSpeed(self):
//         if self.type == EUROPEAN:
//             return self.getBaseSpeed()
//         elif self.type == AFRICAN:
//             return self.getBaseSpeed() - self.getLoadFactor() * self.numberOfCoconuts
//         elif self.type == NORWEGIAN_BLUE:
//             return 0 if self.isNailed else self.getBaseSpeed(self.voltage)
//         else:
//             raise Exception("Should be unreachable")
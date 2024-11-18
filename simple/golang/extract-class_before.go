

type Soldier struct {
	health       int
	damage       int
	weaponStatus int
}

func (soldier *Soldier) getDamage() int {
	return soldier.damage
}

func (soldier *Soldier) attack() {
	// ...
}

// class Soldier:
//     self.health = 0
//     self.damage = 0
//     self.weaponStatus = 0

//     def getDamage(self):
//         # ...

//     def attack(self):
//         # ...
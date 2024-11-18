
type Weapon struct {
	damage       int
	weaponStatus int
}

type Soldier struct {
	health int
	weapon Weapon
}

func (weapon *Weapon) getDamage() int {
	return weapon.damage
}

func (soldier *Soldier) attack() {
	// ...
}

// class Weapon:
//     self.damage = 0
//     self.weaponStatus = 0

//     def getDamage(self):
//         # ...

// class Soldier:
//     health = 0
//     weapon = Weapon()

//     def attack(self):
//         # ...
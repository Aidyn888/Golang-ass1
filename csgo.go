package main

import "fmt"

type rifl struct {
	damage 	int
	name	string
}

func newRifl(name string, damage int) *rifl{
	return &rifl{
		name: name,
		damage: damage,
	}
}

func (s *rifl) useWeapon(c *character) {
	fmt.Printf("shoots the %s with a %s!\n", c.name, s.name)
	c.health -= s.damage
}

type weapon interface {
    useWeapon(opponent *character)
}

type gun struct {
	damage 	int
	name	string
}

func newGun(name string, damage int) *gun{
	return &gun{
		name: name,
		damage: damage,
	}
}

func (b *gun) useWeapon(c *character) {
	fmt.Printf("shoots the %s with a %s!\n", c.name, b.name)
	c.health -= b.damage
}

type character struct {
	health	int
	weapon 	weapon
	name	string
	damage	int
}

func newCharacter(name string) *character {
    return &character{
		name: 	name,
		health: 100,
		damage: 1,
    }
}

func (c *character) equipWeapon(w weapon) {
    c.weapon = w
}

func (c *character) attack(opponent *character) {
	fmt.Printf("The %s ", c.name)
	c.weapon.useWeapon(opponent)
}

func printCharacterStats(c *character) {
	fmt.Printf("The %s has %d health left.\n", c.name, c.health)
}

func main() {
	m416 := newRifl("M416", 25)
	awm := newGun("AWM", 75)
	akm := newRifl("AKM", 45)

	terror := newCharacter("Terrorist")
	terror.equipWeapon(m416)
	ct := newCharacter("Counter-Terrorist")
	ct.equipWeapon(akm)

	printCharacterStats(terror)
	printCharacterStats(ct)

	terror.attack(ct)

	printCharacterStats(ct)

	ct.attack(terror)

	printCharacterStats(terror)

	terror.equipWeapon(awm)
	terror.attack(ct)

	printCharacterStats(ct)
}


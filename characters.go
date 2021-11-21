package main
import (
	"fmt"
	"math/rand"
)
type Player interface {
	Attack( player Player)
	Health () int
	SetHealth(int)
	GetTalent(player Player)
	Defend()
	setDefense(int)
	isAlive() bool
}
type Witch struct {
	Healths int
	Magic int
	Beauty int
	Shield int
}
type Magician struct {
	Healths int
	Magic int
	Age int
}
type Warrior struct {
	Healths int
	Weapon int
	Shield int
}

type Archer struct{
	Healths int
	Strength int
	Numberofarrows int
	Shield int
}

//атака
func (wi *Witch) Attack(player Player) {
	wi.Magic += 15
	wi.Beauty+=2
	damage := wi.Magic + (10 + rand.Intn(11)) - wi.Shield
	fmt.Println("You have dealt ", damage, " damage")
	player.SetHealth(damage)
	player.setDefense(damage)
}
func (wa *Warrior) Attack(player Player) {
	wa.Weapon += 20
	damage := wa.Weapon + (10 + rand.Intn(7)) - wa.Shield
	fmt.Println("You have dealt ", damage, " damage")
	player.SetHealth(damage)
	player.setDefense(damage)
}
func (m * Magician) Attack(player Player) {
	m.Magic += 15
	damage := m.Magic + (10 + rand.Intn(5)) - m.Age
	fmt.Println("You have dealt ", damage, " damage")
	player.SetHealth(damage)
	player.setDefense(damage)
}
func (a * Archer) Attack(player Player) {
	a.Strength += 10
	a.Numberofarrows-=5
	damage := a.Strength + (10 + rand.Intn(15)) - a.Shield
	fmt.Println("You have dealt ", damage, " damage")
	player.SetHealth(damage)
	player.setDefense(damage)
}
//защита
func (wi *Witch) Defend() {
	wi.Magic += 10
	wi.Beauty-=rand.Intn(3)
	wi.Shield = wi.Healths -rand.Intn(15)
}
func (wa *Warrior) Defend() {
	wa.Weapon += 5
	wa.Shield = wa.Healths -rand.Intn(20)
}
func (m *Magician) Defend() {
	m.Magic += 7
	m.Age = m.Age+rand.Intn(5)
}
func (a *Archer) Defend() {
	a.Strength += 10
	a.Numberofarrows-=3
	a.Shield = a.Healths -rand.Intn(12)
}
//дополнительные возможности каждого героя, доступны не всегда
func (wa *Warrior) GetTalent(player Player) {
	if wa.Healths < 50 && wa.Shield<20 {
		fmt.Println("It`s impossible to unlock Superenergy")
		return
	}
	fmt.Println("You have unlocked SUPERENERGY, it means Weapon+10, shield+10+rand")
	wa.Weapon += 10
	wa.Shield += 15+rand.Intn(10)
	damage := wa.Weapon + (10 + rand.Intn(10)) - wa.Shield
	fmt.Println("You have dealt ", damage, " damage")
	player.SetHealth(damage)
	player.setDefense(damage)
}
func (wi *Witch) GetTalent(player Player) {
	if wi.Magic < 35 && wi.Beauty<6 {
		fmt.Println("It`s impossible to unlock Superpower")
		return
	}
	fmt.Println("You have unlocked SUPERPOWER, it means Magic+20, shield+10+rand")
	wi.Magic += 20
	wi.Shield += 10+rand.Intn(10)
	damage := wi.Magic + (10 + rand.Intn(10)) - wi.Shield
	fmt.Println("You have dealt ", damage, " damage")
	player.SetHealth(damage)
	player.setDefense(damage)
}
func (m *Magician) GetTalent(player Player) {
	if m.Magic < 60 && m.Age>70 {
		fmt.Println("It`s impossible to unlock Superpower")
		return
	}
	fmt.Println("You have unlocked SUPERPOWER, it means Magic+10, you become younger")
	m.Magic += 10
	m.Age += m.Age-rand.Intn(10)
	damage := (m.Magic/10) + (10 + rand.Intn(10))
	fmt.Println("You have dealt ", damage, " damage")
	player.SetHealth(damage)
	player.setDefense(damage)
}
func (a *Archer) GetTalent(player Player) {
	if a.Strength < 50 && a.Numberofarrows<30 {
		fmt.Println("It`s impossible to unlock Superpower")
		return
	}
	fmt.Println("You have unlocked SUPERPOWER, it means strength+10, number of arrows +50")
	a.Strength += 10
	a.Shield += 10+rand.Intn(10)
	damage := a.Strength + (10 + rand.Intn(10)) - a.Shield
	fmt.Println("You have dealt ", damage, " damage")
	player.SetHealth(damage)
	player.setDefense(damage)
}
//здоровье и ущерб
func (wi *Witch) Health() int {
	return wi.Healths
}
func (wi *Witch) SetHealth(points int) {
	wi.Healths = wi.Healths - points
}
func (wi *Witch) setDefense(point int) {
	wi.Shield = wi.Shield - point
}
func (wa *Warrior) Health() int {
	return wa.Healths
}
func (wa *Warrior) SetHealth(points int) {
	wa.Healths = wa.Healths - points
}
func (wa *Warrior) setDefense(point int) {
	wa.Shield = wa.Shield - point
}
func (m *Magician) Health() int {
	return m.Healths
}
func (m *Magician) SetHealth(points int) {
	m.Healths = m.Healths - points
}
func (m *Magician) setDefense(point int) {
	m.Age = m.Age + point
}
func (a *Archer) Health() int {
	return a.Healths
}
func (a *Archer) setHealth(points int) {
	a.Healths = a.Healths - points
}
func (a *Archer) setDefense(point int) {
	a.Shield = a.Shield - point
}

func (wa *Warrior) isAlive() bool {
	return wa.Healths > 0
}
func (wi *Witch) isAlive() bool {
	return wi.Healths > 0
}
func (ma *Magician) isAlive() bool {
	return ma.Healths > 0
}
func (a *Archer) isAlive() bool {
	return a.Healths > 0
}
/*witches:=rand.Intn(10)
warriors:= rand.Intn(7)
magicians:=rand.Intn(6)
archers:=rand.Int(8)
sum:=witches+magicians+warriors+archers
if sum/2!=0 {
	magicians+=1
}*/

/*
func main() {
	var choiceOne, choiceTwo string
	fmt.Println("Welcome to Arena, now it`s time to choose your fighters")
	fmt.Println("Choose your first fighter: Wa - warrior, M - magician, Wi-witch, A-archer")
	fmt.Scan(&choiceOne)
	var first, second Player
	switch choiceOne {
	case "Wi":
		first = &Witch{Healths: 100, Magic: 20, Superpower: 10, Shield:30}
	case "Wa":
		first = &Warrior{Healths: 100, Weapon: 20, Superenergy: 10, Shield:30}
	case "M":
		first = &Magician{Healths: 100, Magic: 20, Weapon: 10, Shield:30}
	case "A":
		first = &Archer{Healths: 100, Magic: 20, Weapon: 10, Shield:30}
	}
	fmt.Println("Choose your second fighter: Wa - warrior, M - magician, Wi-witch, A-archer")
	fmt.Scan(&choiceTwo)
	switch choiceTwo {
	case "Wi":
		first = &Witch{Healths: 100, Magic: 20, Superpower: 10, Shield:30}
	case "Wa":
		first = &Warrior{Healths: 100, Weapon: 20, Superenergy: 10, Shield:30}
	case "M":
		first = &Magician{Healths: 100, Magic: 20, Weapon: 10, Shield:30}
	case "A":
		first = &Archer{Healths: 100, Magic: 20, Weapon: 10, Shield:30}
	}
	var choice int
	var whofirst int
	for first.Health() != 0 && second.Health() != 0 {
		fmt.Println("Enter number from 1 to 9")
		fmt.Scan(&whofirst)

		fmt.Println("Choose your action: 1 - attack, 2 - defense, 3 - get talent (cost 10), , 9 - exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			first.Attack(second)
		case 2:
			first.Defend()
		case 3:
			first.GetTalent(second)
		case 9:
			return
		}

		if second.Health() <= 0 {
			fmt.Println("The first one won")
			return
		}
		fmt.Println("Choose your action: 1 - attack, 2 - defense, 3 - first skill (cost 10), 9 - exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			second.Attack(first)
		case 2:
			second.Defend()
		case 3:
			second.FirstSkill(firstHero)
		case 9:
			return
		}
		if first.Health() <= 0 {
			fmt.Println("The second one won")
			return
		}
	}
}*/
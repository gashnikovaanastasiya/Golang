package awesomeProject3

import (
"fmt"
"math/rand"
"time"
)

const count = 256

type AttacksPool struct{
	name string
	dmg float32
	cost float32
}
type Warrior struct {
	name string
	Healths float32
	Weapon int
	Shield int
	attackpool []AttacksPool
	unitType string
}
type Witch struct {
	name string
	Healths float32
	Magic int
	Beauty int
	Shield int
	attackpool []AttacksPool
	unitType string
}

type Archer struct{
	name string
	Healths float32
	Numberofarrows int
	Shield int
	attackpool []AttacksPool
	unitType string
}
type Mage struct {
	name string
	Healths float32
	Magic int
	Age int
	attackpool []AttacksPool
	unitType string
}

type Player interface{
	Attack(Player)
	ChooseSkill() int
	setHealth(float32)
	getHealth() float32
	getName() string
	printInfo()
}

func(mage *Mage) setHealth(num float32){
	mage.Healths = mage.Healths - num
}
func(war *Warrior) setHealth(num float32){
	war.Healths = war.Healths - num
}
func(ar *Archer) setHealth(num float32){
	ar.Healths = ar.Healths - num
}
func(wi *Witch) setHealth(num float32){
	wi.Healths = wi.Healths - num
}

func(mage *Mage) getHealth() float32{
	return mage.Healths
}
func(war *Warrior) getHealth() float32{
	return war.Healths
}
func(wi *Witch) getHealth() float32{
	return wi.Healths
}
func(ar *Archer) getHealth() float32{
	return ar.Healths
}

func(mage *Mage) getName() string{
	return mage.name
}
func(war *Warrior) getName() string{
	return war.name
}
func(wi *Witch) getName() string{
	return wi.name
}
func(ar *Archer) getName() string{
	return ar.name
}
func(mage *Mage) Attack(player Player){
	if (mage.Healths >= mage.attackpool[mage.ChooseSkill()].cost) && (mage.Age<=100) {
		damage := mage.attackpool[mage.ChooseSkill()].dmg
		player.setHealth(damage)
		mage.Healths -= mage.attackpool[mage.ChooseSkill()].cost
		fmt.Printf(mage.name + " done succ attack and dealed %.2f dmg by %s ," + player.getName() + " have %.2f hp\n",mage.attackpool[mage.ChooseSkill()].dmg,mage.attackpool[mage.ChooseSkill()].name,player.getHealth())
	}else{
		fmt.Print("You have dead\n")
	}
	mage.Healths += 100
	mage.Age-=5
}
func(war *Warrior) Attack(player Player){
	if (war.Healths >= war.attackpool[war.ChooseSkill()].cost) &&(war.Shield>5) {
		damage := war.attackpool[war.ChooseSkill()].dmg
		player.setHealth(damage)
		war.Weapon +=10
		war.Healths -= war.attackpool[war.ChooseSkill()].cost
		war.Shield+=5
		fmt.Printf(war.name + " done succ attack and dealed %.2f dmg by %s , " + player.getName() +" have %.2f hp\n",war.attackpool[war.ChooseSkill()].dmg,war.attackpool[war.ChooseSkill()].name,player.getHealth())
	}else{
		fmt.Print("You have dead\n")
	}
	war.Shield += 1
}
func(wi *Witch) Attack(player Player){
	if (wi.Healths >= wi.attackpool[wi.ChooseSkill()].cost) &&(wi.Shield>10) {
		damage := wi.attackpool[wi.ChooseSkill()].dmg
		player.setHealth(damage)
		wi.Healths -= wi.attackpool[wi.ChooseSkill()].cost
		wi.Shield+=5
		wi.Beauty+=5
		fmt.Printf(wi.name + " done succ attack and dealed %.2f dmg by %s , " + player.getName() +" have %.2f hp\n",wi.attackpool[wi.ChooseSkill()].dmg,wi.attackpool[wi.ChooseSkill()].name,player.getHealth())
	}else{
		fmt.Print("You have dead\n")
	}
	//wi.Shield += 2
}
func(ar *Archer) Attack(player Player){
	if (ar.Healths >= ar.attackpool[ar.ChooseSkill()].cost) &&(ar.Shield>5) {
		damage := ar.attackpool[ar.ChooseSkill()].dmg
		player.setHealth(damage)
		ar.Healths -= ar.attackpool[ar.ChooseSkill()].cost
		ar.Numberofarrows+=5
		fmt.Printf(ar.name + " done succ attack and dealed %.2f dmg by %s , " + player.getName() +" have %.2f hp\n",ar.attackpool[ar.ChooseSkill()].dmg,ar.attackpool[ar.ChooseSkill()].name,player.getHealth())
	}else{
		fmt.Print("You have dead\n")
	}
	ar.Shield += 3
}

//выбор способностей
func(mage Mage) ChooseSkill() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(mage.attackpool))
}
func(war Warrior) ChooseSkill() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(war.attackpool))
}
func(wi Witch) ChooseSkill() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(wi.attackpool))
}
func(ar Archer) ChooseSkill() int{
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(ar.attackpool))
}

func setWarrior(war* Warrior) *Warrior{
	time.Sleep(10)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	names := []string{"Max","Mike","Alex","Ron"}
	warrior := *war
	warrior.name = names[num]
	warrior.Healths = 100
	warrior.Weapon = 50
	warrior.Shield = 1 + 5*rand.Intn(4)
	warrior.attackpool = append(warrior.attackpool, AttacksPool{"AttackWithSuperpower",10,8})
	warrior.attackpool = append(warrior.attackpool, AttacksPool{"UsualWeapon",6,6})
	warrior.attackpool = append(warrior.attackpool, AttacksPool{"WeaponToDefense",2,5})
	return &warrior
}
func setMage(mg* Mage) *Mage{
	time.Sleep(10)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	names := []string{"Mag1","Mag2","Mag3","Mag4","Mag5"}
	mage := *mg
	mage.name = names[num]
	mage.Healths = 70
	mage.Magic = 50
	mage.Age = 50
	mage.attackpool = append(mage.attackpool, AttacksPool{"SuperMagic",10,9})
	mage.attackpool = append(mage.attackpool, AttacksPool{"MagicForAttack",9,8})

	mage.attackpool = append(mage.attackpool, AttacksPool{"MagicForDefense",1,3})
	return &mage
}
func setArcher(ar* Archer) *Archer{
	time.Sleep(10)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	names := []string{"Archer1","Archer2","Archer3","Archer4","Archer5"}
	arch := *ar
	arch.name = names[num]
	arch.Healths = 100
	arch.Numberofarrows = 100
	arch.Shield = 25 - 7*rand.Intn(2)
	arch.attackpool = append(arch.attackpool, AttacksPool{"StrengthsForArcher",10,9})
	arch.attackpool = append(arch.attackpool, AttacksPool{"+20Arrows",9,8})
	arch.attackpool = append(arch.attackpool, AttacksPool{"ShieldForDefence",1,4})
	return &arch
}
func setWitch(wi* Witch) *Witch{
	time.Sleep(10)
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(10)
	names := []string{"Kendall","Kylie","Kim","Khloe","Kourtney"}
	witch := *wi
	witch.name = names[num]
	witch.Healths = 100
	witch.Magic = 100
	witch.Shield = 11 + rand.Intn(10)
	witch.attackpool = append(witch.attackpool, AttacksPool{"SuperMagic",10,9})
	witch.attackpool = append(witch.attackpool, AttacksPool{"WitchsMagic",9,8})
	witch.attackpool = append(witch.attackpool, AttacksPool{"MagForDefense",1,0})
	return &witch
}

//создание персонажей
func createWarrior() *Warrior{
	warrior := Warrior{}
	warrior.unitType = "Warrior"
	return setWarrior(&warrior)
}
func createMage() *Mage{
	mage := Mage{}
	mage.unitType = "Mage"
	return setMage(&mage)
}
func createWitch() *Witch{
	witch := Witch{}
	witch.unitType = "Witch"
	return setWitch(&witch)
}
func createArcher() *Archer{
	archer := Archer{}
	archer.unitType = "Archer"
	return setArcher(&archer)
}
//вывод инофрмации о герое
func (war *Warrior) printInfo(){
	fmt.Printf(" : %s , %s have %.2f hp and %.2f def \n",war.name,war.unitType,war.Healths,war.Weapon,war.Shield)
}
func (mage *Mage) printInfo(){
	fmt.Printf(" : %s , %s have %.2f hp and %.2f def \n", mage.name,mage.unitType,mage.Healths,mage.Magic, mage.Age)
}
func (wi *Witch) printInfo(){
	fmt.Printf(" : %s , %s have %.2f hp and %.2f def \n", wi.name,wi.unitType,wi.Healths,wi.Magic, wi.Beauty)
}
func (ar *Archer) printInfo(){
	fmt.Printf(" : %s , %s have %.2f hp and %.2f def \n", ar.name,ar.unitType,ar.Healths,ar.Numberofarrows, ar.Shield)
}


func fight(p1 Player,p2 Player)  int{
	p1.Attack(p2)
	if p2.getHealth() <= 0 {

		fmt.Println(p2.getName() + " dead")
		return 1
	} else {
		p2.Attack(p1)
		if p1.getHealth() <= 0 {
			fmt.Println(p1.getName() + " dead")
			return 2
		}
	}
	return 0
}

func theLastOneRemained(players []Player) bool{
	if(len(players) == 1){
		return true
	} else{
		return false
	}
}
func isCrit()bool{
	if(rand.Intn(5) == 0){
		return true
	} else{
		return false
	}
}
func printName(player Player) string{
	return player.getName()
}

func printInfo(player Player){
	player.printInfo()
}
func checkForEven(count int) bool{
	if count%2 == 0{
		return true
	} else{
		return false
	}
}

func battleTillTheEnd(i int,players []Player,message chan Player){
	personNum1 := 2*i
	personNum2 := 2*i+1
	winner := 0
	for ;winner == 0; {
		winner = fight(players[personNum1], players[personNum2])
	}
	if (winner == 1) {
		message <- players[personNum1]
	} else{
		message <- players[personNum2]
	}
}
func recovery(players []Player) []Player{
	for i:= range(players) {
		players[i].setHealth(-50)
		//heroes[i].setAdditionalStats(-50)
	}
	return players
}
func main(){
	if(!checkForEven(count)){
		print("Error! Not even count of players")
		return
	}
	startTime := time.Now()
	message := make(chan Player)
	players := make([]Player,count)
	tempplayers := make([]Player,count)
	rand.Seed(time.Now().UnixNano())
	for i := range players{
		num := rand.Intn(4)
		switch num{
		case 0:
			temp:= createWarrior()
			players[i] = temp
		case 1:
			temp:= createMage()
			players[i] = temp
		case 2:
			temp:= createWitch()
			players[i] = temp
		case 3:
			temp:= createArcher()
			players[i] = temp
		}
	}
	println("Players : ")
	for i:= range(players){
		printInfo(players[i])
	}
	for ;len(players)!= 0 &&!theLastOneRemained(players);{
		time.Sleep(1*time.Millisecond)
		tempplayers =nil
		heroesLength := len(players)
		for i:=0;i < heroesLength/2;i++ {
			go battleTillTheEnd(i, players, message)
		}
		for i:=0;i<heroesLength/2;i++ {
			tempplayers = append(tempplayers, <-message)
		}
		if(heroesLength!=2) {
			recovery(tempplayers)
		}
		players = tempplayers
	}
	print(players[0].getName())
	println(" survived")
	printInfo(players[0])
	timeHasPassed := time.Since(startTime)
	print("time has passed since launch = ");print(timeHasPassed.Seconds()); print(" seconds")
}

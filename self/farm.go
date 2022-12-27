package self

import (
	"fmt"
	"math/rand"
	"time"
)

type honeyBee struct {
	name string
}

func (b honeyBee) String() string {
	return b.name
}
func (hb honeyBee) move() string {
	switch rand.Intn(2) {
	case 0:
		return "viz viz"
	default:
		return "uchish va damolish"
	}
}
func (hb honeyBee) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "nektar"
	default:
		return "gul chang"
	}
}

type gopher struct {
	name string
}

func (g gopher) String() string {
	return g.name
}
func (g gopher) move() string {
	switch rand.Intn(2) {
	case 0:
		return "yurish va hududni o'rganish"
	default:
		return "yashiringan"
	}
}
func (g gopher) eat() string {
	switch rand.Intn(5) {
	case 0:
		return "sabzi"
	case 1:
		return "salat"
	case 2:
		return "makka jo'xori"
	case 3:
		return "rediska"
	default:
		return "ildiz"
	}
}

type animal interface {
	eat() string
	move() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v. \n", a, a.move())
	default :
		fmt.Printf("%v ovqatlanmoqda %v. \n", a, a.eat())
	}

}
const sunrise, sunset = 8, 18
func Farm(){
	rand.Seed(time.Now().UnixNano())
animal := []animal{
	honeyBee{name: "Asal ari"},
	gopher{name: "Go"},
}
var sol, hour int
for {
	fmt.Printf("%d : 00", hour)
	if hour < sunrise || hour >= sunset{
		fmt.Println("Hayvonlar uyquda. ")
	}else{
		i := rand.Intn(len(animal))
		step(animal[i])
	}
	time.Sleep(500 * time.Millisecond)
	hour ++
	if hour >= 24{
		hour = 0
		sol ++
		if sol >= 3{
			break
		}
	}
}
}

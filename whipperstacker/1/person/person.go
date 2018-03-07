package person

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Person is a simple actor structure
type Person struct {
	name      string
	fireAlarm chan<- bool
	exit      chan<- bool
	wg        *sync.WaitGroup
}

// New returns a new person with given name
func New(name string, fireAlarm, exit chan<- bool, wg *sync.WaitGroup) *Person {
	return &Person{
		name:      name,
		fireAlarm: fireAlarm,
		exit:      exit,
		wg:        wg,
	}
}

// ready is the first step of routine
func (p *Person) ready() int {
	wait := rand.Intn(30) + 60 // [60,90]
	t := time.After(time.Duration(wait) * time.Second)
	<-t
	return wait
}

// shoes simulate the action of the person putting on shoes
func (p *Person) shoes() int {
	wait := rand.Intn(10) + 35 // [35,45]
	t := time.After(time.Duration(wait) * time.Second)
	<-t
	return wait
}

// Live is the morning routine of the person
func (p *Person) Live() {
	defer p.wg.Done()

	fmt.Println(p.name, "started getting ready")
	wait := p.ready()
	fmt.Println(p.name, "spent", wait, "seconds getting ready")
	p.fireAlarm <- true

	fmt.Println(p.name, "started putting on shoes")
	wait = p.shoes()
	fmt.Println(p.name, "spent", wait, "seconds putting on shoes")
	p.exit <- true
}

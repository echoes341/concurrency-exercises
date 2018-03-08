package main

import (
	"sync"
)

func main() {
	tourists := make(chan tourist)
	wg := sync.WaitGroup{}
	go pcPool(tourists)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		tourists <- tourist{uint(i), &wg}
	}

	wg.Wait()
}

func pcPool(personPool chan tourist) {
	personBuffer := make(chan tourist, 8)

	for i := 0; i < 8; i++ {
		go computer(personBuffer)
	}

	for {
		person := <-personPool
		personBuffer <- person
	}
}

func computer(person chan tourist) {
	for {
		p := <-person
		p.Hello()
	}
}

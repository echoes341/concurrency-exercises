package main

import (
	"sync"

	"github.com/echoes341/concurrency-exercises/whipperstacker/1/person"
)

func main() {
	wg := &sync.WaitGroup{}
	alarmTrg := make(chan bool, 2)
	doorTrg := make(chan bool, 2)

	alice := person.New("Alice", alarmTrg, doorTrg)
	bob := person.New("Bob", alarmTrg, doorTrg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		alice.Live()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		bob.Live()
	}()

	wg.Add(1)
	go alarm(alarmTrg, wg)

	wg.Add(1)
	go door(doorTrg, wg)

	wg.Wait()
}

func alarm(trigger <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	<-trigger
	<-trigger

}

func door(trigger <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
}

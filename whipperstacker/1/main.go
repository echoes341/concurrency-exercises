package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/echoes341/concurrency-exercises/whipperstacker/1/person"
)

func main() {
	wg := &sync.WaitGroup{}
	alarmTrg := make(chan bool, 2)
	doorTrg := make(chan bool, 2)

	alice := person.New("Alice", alarmTrg, doorTrg, wg)
	bob := person.New("Bob", alarmTrg, doorTrg, wg)

	fmt.Println("Let's go for a walk!")

	wg.Add(4)
	go alice.Live()
	go bob.Live()
	go alarm(alarmTrg, wg)
	go door(doorTrg, wg)

	wg.Wait()
}

func alarm(trigger <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	<-trigger
	<-trigger

	fmt.Println("Arming alarm.")

	t := time.After(60 * time.Second)
	fmt.Println("Alarm is counting down.")

	<-t
	fmt.Println("Alarm is armed.")
}

func door(trigger <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	<-trigger
	<-trigger

	fmt.Println("Exiting and locking the door.")
}

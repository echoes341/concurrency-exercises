package main

import (
	"sync"

	"github.com/echoes341/concurrency-exercises/whipperstacker/3/tourist"
)

var wg = sync.WaitGroup{}

func main() {
	tourists := [25]*tourist.Tourist{}
	for k := range tourists {
		tourists[k] = tourist.New(k)
		wg.Add(1)
		go tourists[k].Run(&wg)
	}
	wg.Wait()
}

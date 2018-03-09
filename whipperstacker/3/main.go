package main

import (
	"fmt"
	"sync"
)

type status int

const (
	ONLINE = status(iota)
	WAITING
)

var wg = sync.WaitGroup{}

func main() {
	tourists := [25]tourist{}
	for k := range tourists {
		tourists[k].n = k
		wg.Add(1)
		go tourists[k].Run()
	}
	wg.Wait()
}

type tourist struct {
	n int
	status
}

func (t *tourist) Run() {
	defer wg.Done()
	fmt.Println(t.n)
}

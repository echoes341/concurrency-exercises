package main

import (
	"fmt"
	"sync"
	"time"
)

type tourist struct {
	number uint
	wg     *sync.WaitGroup
}

func (p tourist) Hello() {
	fmt.Println("Hello", p.number)
	time.Sleep(2 * time.Second)
	p.wg.Done()
}

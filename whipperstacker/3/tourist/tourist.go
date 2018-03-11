package tourist

import (
	"fmt"
	"sync"
	"time"

	"math/rand"

	"github.com/echoes341/concurrency-exercises/whipperstacker/3/pool"
)

// Tourist is our actor
type Tourist struct {
	n        int
	waitTime int
	wg       sync.WaitGroup
}

func init() {
	rand.Seed(time.Now().Unix())
}

// New returns a new tourist
func New(n int) *Tourist {
	return &Tourist{
		n:        n,
		waitTime: rand.Intn(8) + 5,
		wg:       sync.WaitGroup{},
	}
}

// Run tourist behaviour
func (t *Tourist) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	pcPool, wait := pool.Computer()
	if wait { // computer is not ready
		fmt.Println("Tourist", t.n, "waiting for turn.")
	}
	// wait till it receives a channel to pc available
	pc := <-pcPool
	fmt.Println("Tourist", t.n, "is online")
	// now I can send task
	t.wg.Add(1)
	pc <- t
	t.wg.Wait()
}

// Exec is tourist task
func (t *Tourist) Exec() {
	wait := time.Duration(t.waitTime) * time.Second
	time.Sleep(wait)
}

// Finished is closing task after Exec
func (t *Tourist) Finished() {
	fmt.Println("Tourist", t.n, "is done, having spent", t.waitTime, "minutes online.")
	t.wg.Done()
}

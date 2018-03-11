package pool

import (
	"sync"
)

const numPc = 8

var pcBuff chan chan Executable

var availables struct {
	sync.Mutex
	av int
}

func init() {
	pcBuff = make(chan chan Executable, numPc)

	availables = struct {
		sync.Mutex
		av int
	}{
		Mutex: sync.Mutex{},
		av:    numPc,
	}

	for i := 0; i < numPc; i++ {
		go pc()
	}
}

func pc() {
	fnChan := make(chan Executable)
	pcBuff <- fnChan // first case: pc is available

	for {
		task := <-fnChan
		task.Exec() // execute function
		task.Finished()

		// pc is available again
		incAvailables()
		pcBuff <- fnChan
	}
}

func incAvailables() {
	availables.Lock()
	defer availables.Unlock()
	availables.av++
}

// Computer gets the main pc pool and wait status
func Computer() (pool chan chan Executable, wait bool) {
	availables.Lock()
	defer availables.Unlock()

	if availables.av == 0 {
		wait = true
	} else {
		availables.av--
	}

	return pcBuff, wait
}

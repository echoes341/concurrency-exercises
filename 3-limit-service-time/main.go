//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		process()
		return true
	}

	var status = make(chan bool)
	tmr := time.NewTimer(10 * time.Second)
	go execute(process, status)
	select {
	case <-tmr.C:
		return false
	case <-status:
		return true
	}

}

func execute(process func(), finished chan<- bool) {
	process()
	finished <- true
}

func main() {
	RunMockServer()
}

package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	sign := make(chan os.Signal)
	signal.Notify(sign, os.Interrupt)

	for {
		s := <-sign
		fmt.Println("Received signal!")
		fmt.Println(s)
	}
}

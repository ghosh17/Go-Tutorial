package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var Balance float64 = 50.00

func main() {
	//Opening & Closing = Open and close channel
	// deposite account <- $$
	//Withdrawal $$ <- account

	wg.Add(2)
	c := make(chan float64, 50) //buffer
	dflag := 0
	d := 123406.78
	w := 0.00
	if d > 0 {
		dflag = 1
	}
	money := d - w

	go func() {
		for i := range c {
			time.Sleep(time.Duration(3 * time.Second))
			fmt.Printf("You got $%0.2f M's in your bank account", i)
			wg.Done()
		}

	}()

	go func() {

		if dflag == 1 {
			c <- (Balance + money)
		} else {
			c <- (Balance - money)
		}

		wg.Done()
	}()

	defer close(c) //Shut down Bank account

	wg.Wait()

}

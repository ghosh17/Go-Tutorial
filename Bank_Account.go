package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var wg2 = sync.WaitGroup{}

var Balance float64 = 100000.00

func main() {
	//Opening & Closing = Open and close channel
	// deposite account <- $$
	//Withdrawal $$ <- account

	wg.Add(2)
	c := make(chan float64)

	deposite := 23456.79
	withdraw := 0.01

	go func() {
		time.Sleep(time.Duration(5 * time.Second))
		for i := range c {
			time.Sleep(time.Duration(5 * time.Second))
			fmt.Printf("You got $%0.2f M's in your bank account", i)
			wg.Done()
		}

	}()

	go func() {

		wg2.Add(2)

		go Deposite(c, deposite)
		//c <- (Balance + money)

		go Withdraw(c, withdraw)
		//c <- (Balance + money)

		wg2.Wait()
		wg.Done()
	}()

	defer close(c) //Shut down Bank account

	wg.Wait()

}

func Deposite(c chan float64, deposite float64) {
	fmt.Printf("Depositing: %0.2f\n", deposite)
	Balance = Balance + deposite
	c <- (Balance)
	wg2.Done()
}

func Withdraw(c chan float64, withdraw float64) {
	fmt.Printf("Withdrawing: %0.2f\n", withdraw)
	Balance = Balance - withdraw
	c <- (Balance)
	wg2.Done()
}

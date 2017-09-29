//Recursion in Go
package main

import "fmt"

func main() {
	//Fib series

	for i := 0; i < 10; i++ {
		fmt.Printf(" %d ", fib(i))
	}

}

func fib(num int) int {
	if num == 0 {
		return 0
	}
	if num == 1 {
		return 1
	}
	return fib(num-1) + fib(num-2)
}

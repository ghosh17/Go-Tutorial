//Fuction, if and for loop in GO
package main

import "fmt"

/*
Main Function
*/
func main() {
	fmt.Print("\nFunctions, conditionals and loops in GO\n")
	var a, b int = 5, 6
	fmt.Print("a=", a, "b=", b, "\n")
	c, d := swap(a, b)
	fmt.Print("a=", c, "b=", d, "\n")

	fmt.Print("\n Now doing Loops \n")

	count := loop_func(a)

	fmt.Print("The value of count is ", count, "\n")
}

func swap(x, y int) (int, int) {
	return y, x
}

func loop_func(val int) int {

	for i := 0; i < 100; i++ {
		val = val + 1
	}
	return val
}

//Pointer in GO
package main

import "fmt"

func main() {

	var ptr *int //integer Pointer
	var a int = 5

	ptr = &a
	fmt.Printf("Address of variable %x \n", &a)
	fmt.Printf("ptr pointing to location %x \n", ptr)
	fmt.Printf("Value/data that ptr is pointing to is %d \n", *ptr)

	//Null Nil Pointers in Go

	var uninitialized_ptr *int

	fmt.Printf("The Value of an uninitialized pointer in go is : %d \n", uninitialized_ptr)

}

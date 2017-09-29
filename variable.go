package main

import "fmt"

func main() {
	fmt.Print("This is a program to learn about variable data types in GO \n")

	var i, j, k int //This is initialized to int
	var f float64   //float is a float64 variable
	fmt.Print("The value of i = %d and the variable type is = %T \n", i, i)

	//Mixed variable declaration

	var a, b, c, foo = 3, 4.00, 5.0, "Alienation of Labour" //a=>int b,c=>float foo=>string

	fmt.Print("The value of foo is")
	fmt.Println(foo)
	fmt.Print(" and the variable type is = %T \n")

	//Dynamic type declaration

	p := 42
	fmt.Print("The value of p is")
	fmt.Println(p)
	fmt.Print("and the variable type is = %T \n", p)

	//LValue=Expressions that refers to a memory location
	//RValue=Expressions that refer to the data stored at a perticular location

}

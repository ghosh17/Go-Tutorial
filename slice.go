//Slices: Arrays of variable size
package main

import "fmt"

func main() {

	var numbers []int
	//or
	var number = make([]int, 3, 5) //Make a slice of int with len=0(number of elements preset)
	//and cap=5(how many elements can be accomodated)

	numbers = number

	printSlice(numbers)
	/*
	   //Subslicing
	*/

	num := []int{0, 1, 2, 3, 4, 5}
	printSlice(num)

	num1 := num[2:4] //Slice num and print 2->4
	printSlice(num1)

	/*
	  //Append and Copy functions
	*/

	//append adds elements to the slice

	numbers = append(numbers, 0)
	printSlice(numbers)
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 6, 7)
	printSlice(numbers)
	numbers2 := make([]int, len(numbers), (cap(numbers))*4)
	copy(numbers2, numbers)
	printSlice(numbers2)

}

func printSlice(x []int) {
	fmt.Printf("\n****\n The len = %d & the Capacity is = %d\n slice = %v", len(x), cap(x), x)
}

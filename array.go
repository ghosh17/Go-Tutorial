//Array in GO
package main

import "fmt"

func main() {
	//Declare an Array

	var arr [10]int //Declared an array of size 10 of type interger

	var i, j int

	for i = 0; i < 10; i++ {
		arr[i] = i
	}

	for j = 0; j < 10; j++ {
		fmt.Printf("Position arr[%d] =  %d \n", j, arr[j])
	}

	//Multidimentional arrays in GO
	var m_arr = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var row int = 2
	var col int = 3
	var avg float64
	avg = Get_Avg(m_arr, row, col)
	fmt.Printf("\nAverage %f \n", avg)
}

//Passing array to functions
func Get_Avg(arr [2][3]int, row int, col int) float64 {
	var i, j int
	var sum int = 0
	var avg float64
	for i = 0; i < row; i++ {
		for j = 0; j < col; j++ {
			sum += arr[i][j]
		}
	}
	avg = float64(sum / (6))
	return avg

}

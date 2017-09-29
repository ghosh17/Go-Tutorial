//package binarysearch

//Implementing Binary seach in GO
package main

import "fmt"
import "errors"

func binary_search(arr []int, low int, high int, ele int) int {

	var mid int

	if high < low {
		errors.New("Array non existant.")
	}

	mid = (high + low) / 2

	if arr[mid] == ele { //If the mid is the key
		return mid
	} else if ele > arr[mid] { //If the key is present at a value greater than the mid
		mid = binary_search(arr, mid+1, high, ele)
	} else { //If the key is present at a value lesser than the mid
		mid = binary_search(arr, low, mid-1, ele)
	}
	return mid
}

func main() {
	array := []int{0, 1, 3, 5, 7, 8, 9, 21, 100, 233, 344, 556, 666, 777, 888, 999}

	index := binary_search(array, 0, len(array), 21) //21 savage

	fmt.Printf("@ arr[%d] ", index)
}

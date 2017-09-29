//Mapping function in Go
package main

import "fmt"

func main() {
	//create a Map
	var Map map[int]string
	Map = make(map[int]string)
	Map[1] = "Albert"
	Map[2] = "Carl"
	Map[3] = "Satre"
	Map[4] = "Shelly"
	Map[5] = "Kikergard"

	//print the map
	for number := range Map {
		fmt.Printf("Number : %d Name : %s \n", number, Map[number])
	}
	fmt.Print("\n*****************\n")
	//Check for stuff in a Map
	name, ok := Map[1]
	if ok {
		fmt.Printf("Name : %s \n", name)
	} else {
		fmt.Print("Subject not present \n")
	}

	fmt.Print("\n*****************\n")

	//Delete function in Map

	delete(Map, 5)
	for number := range Map {
		fmt.Printf("Number : %d Name : %s \n", number, Map[number])
	}

}

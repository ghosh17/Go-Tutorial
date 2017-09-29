//Simplest interface to understand interfaces

package main

import "fmt"

type Interface interface {
	Print()
}

type Person struct {
	name string
}

func (p Person) Print() {
	fmt.Printf("Hello %s", p.name)
}

func main() {
	var i Interface = Person{"Sisyphus"}
	i.Print()

}

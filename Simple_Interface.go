//Simplest interface to understand interfaces

package main

type Interface interface {
	Print()
}

type Person struct {
	name string
}

func (p Person) Print() string {
	return "Hello " + p.name
}

func main() {
	var name Person = Person{"Sisyphus"}
	str := name.Print()
	println(str)

}

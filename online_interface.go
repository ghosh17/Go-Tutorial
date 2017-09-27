//Interface for retailers
package main

import "fmt"

func main(){
  //I want to be able to greet my onine visitors
  c1:= customer{"21", "Savage"}
  u1:= user{"kod", "Bla"}
  fmt.Printf("%s\n", c1.greet())
  fmt.Printf("%s\n", u1.greet())
}
//Interface
type namer interface{
  greet() string
}

func (c1 customer) greet()(string){
  return "Hello Customer"
}

func (u1 user) greet()(string){
  return "Hello User"
}

type customer struct{
  First_name string
  Last_name string
}

type user struct{
  First_name string
  Last_name string
}

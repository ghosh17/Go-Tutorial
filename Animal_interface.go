//Interface for Animal
//ANything that speaks is an Animal

package main

import "fmt"

func main(){
  Cade := Cat{}
  Doggo := Dog{}
//  Barn_doggo := Horse{}
  c:=Cade.Speak()
  fmt.Printf("Cat goes %s \n",c)
  fmt.Printf("Dog goes %s \n",Doggo.Speak() )
  //fmt.Printf("Cat goes %s",Speak(Doggo))
}

//Struct for different animals

type Dog struct{

}

type Cat struct{

}

type Horse struct{

}

type Animal interface{
  Speak() string
}

func (Cade Cat) Speak() (string){
  return "Meow"
}

func (Doggo Dog) Speak() (string){
  return "Bark"
}

func (Horsie Horse) Speak() (string){
  return "Neigh"
}

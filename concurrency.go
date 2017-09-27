//Program to learn concurrency in Go
package main

import "fmt"

func main(){
//Without concurrency everything is sequencial
//Have two functions, with out concurrency will execute sequencially one after other
/*
***Concurrency                       VS Parallelism
Doing one task at a time but            Doing Multiple tasks at the
switching between tasks                 same time
eg:talking and eating.                  eg: Talking and driving
Can't do both at exactly same
time
*/


  foo()
  bar()

}

func foo(){
  for i:=0;i<10;i++{
    fmt.Printf("Foo %d \n",i)
  }
}

func bar(){
  for i:=0;i<10;i++{
    fmt.Printf("Bar %d \n",i)
  }
}

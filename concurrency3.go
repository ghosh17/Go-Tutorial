//Program to learn concurrency in Go
package main

import "fmt"

import "sync" //For the concurrency

import "time"

var wg sync.WaitGroup //Wait group library from go

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

//Now we are adding time delays
  wg.Add(2) //Two wait groups added. It will wait to 2 wait groups to execute before exiting main
  go foo()
  go bar()
  wg.Wait()


}

func foo(){
  for i:=0;i<10;i++{
    fmt.Printf("Foo %d \n",i)
    time.Sleep(time.Duration(3*time.Millisecond)) //Each time foo prints sleep for 3 ms
  }
  wg.Done()
}

func bar(){
  for i:=0;i<10;i++{
    fmt.Printf("Bar %d \n",i)
    time.Sleep(time.Duration(10*time.Millisecond))//Each time Bar Prints sleep for 10 ms
  }
  wg.Done()
}

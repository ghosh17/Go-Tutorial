//channels
/*
Prevents Race conditions. Async Data. 2 way channel. Channels can store data

*/
package main

import (
  "fmt"
  "sync"
)

var wg sync.WaitGroup
func main(){
  ch := make(chan int)
  wg.Add(2)
  //**The Transmitter & Receiver work simultaniiously
  //Receiver
  go func(){  //This is called a Anonymous function. Can be called only once from only one place.
    i:=<-ch//read value in channel and store it in i
    fmt.Println(i)//print the value of channel
    wg.Done()
  }()
  //Transmitter
  go func(){
    ch<-42 //Put the value of 42 in channel
    wg.Done()
  }()

  wg.Wait()

}

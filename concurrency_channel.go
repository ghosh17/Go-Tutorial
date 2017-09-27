//Program to learn concurrency in Go
//Channels are like Pipelines
package main

import ("fmt"
        "sync" //For the concurrency
        "time"//for waiting
        "math/rand"//We gonna sleep a random interval
        "sync/atomic"
      )

var wg sync.WaitGroup //Wait group library from go
var mutex sync.Mutex //Locks
//Create these two channels
var c = make(chan int)
var done = make(chan bool)
var x int64

var counter int
func main(){
//Without concurrency everything is sequencial
//Have two functions, with out concurrency will execute sequencially one after other


//Now we are adding time delays
  go incrementor("Foo:")
  go incrementor("Bar:")
  go puller()
  <-done
  fmt.Printf("Final Counter: %d", counter)

}

//Data Race happens when we do this without mutex locks. R/W problem.
//Without mutex: The counter dosen't increase as it should. Foo accesses the counter at 0 as well as bar and they both write 1 to the counter when the counter should be 2 now
//With mutex: We lock one function so that it can finish the job before the next one takes ove
func incrementor(str string){
  for i:=0;i<20;i++{
    time.Sleep(time.Duration(rand.Intn(3))*time.Millisecond)//Sleep a random interval around 3 ms
    c<-1 //Putting something on a channel
    fmt.Println(str,i) //Print the function name and the current interation of it
    if(i==19){
      atomic.AddInt64(&x,1) //Check to see if one of the func has reached end.
      fmt.Println("XXXXXX", x)
    }
    //Check to see if both the functions have reached end.
    if(atomic.LoadInt64(&x)==2){
      close(c) //Close the channel C
    }

    }
}

func puller(){
  for{
    i, more:=<-c
    if(more){
      counter+=i// I
      fmt.Println("Counter: ",counter)
    }else{
      done<-true
      close(done)
      return
    }
  }
}

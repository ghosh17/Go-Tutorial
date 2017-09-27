//Program to learn concurrency in Go
package main

import ("fmt"
        "sync" //For the concurrency
        "time"//for waiting
        "math/rand"//We gonna sleep a random interval
      )

var wg sync.WaitGroup //Wait group library from go
var mutex sync.Mutex //Locks

var counter int
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
  go incrementor("Foo:")
  go incrementor("Bar:")
  wg.Wait()

  fmt.Printf("Final Counter: %d", counter)

}

//Data Race happens when we do this without mutex locks. R/W problem.
//Without mutex: The counter dosen't increase as it should. Foo accesses the counter at 0 as well as bar and they both write 1 to the counter when the counter should be 2 now
//With mutex: We lock one function so that it can finish the job before the next one takes ove
func incrementor(str string){
  for i:=0;i<20;i++{
    time.Sleep(time.Duration(rand.Intn(3))*time.Millisecond)//Sleep a random interval around 3 ms
    mutex.Lock()
    counter++
    fmt.Printf("%s %d Counter: %d\n",str,i,counter)
    mutex.Unlock()
}
wg.Done()
}

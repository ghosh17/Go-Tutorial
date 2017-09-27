//Implement sum of an array using channels
package main

import "fmt"
//import "sync"

func main(){
  ch := make(chan int)//Create a channel of size 50
  Slice := []int{1,2,3,4,5,6,7,8,8,10}
  length:=len(Slice)
  half_length:=len(Slice)/2
  go sum(Slice[0:half_length],ch)
  go sum(Slice[half_length:length],ch)
  a, b:= <-ch, <-ch
  fmt.Printf("The value of two halves are %d + %d = %d \n",a,b,a+b)
}

func sum (arr []int,ch chan int){
  sum := 0
  for _, value := range arr{
    sum+=value
  }
  ch<-sum

}

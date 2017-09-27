//Error handling in Go
package main

import "fmt"
import "math"
import "errors"

func main(){
  //Return a error value when a negative number is passed to a square root function
  var num float64
  num=-3
  result, err:= Sqrt(num)
  if(err==nil){
    fmt.Println("\nThe SQRT is ",result)
  }else{
    fmt.Println(err)
  }
}

func Sqrt(value float64) (float64, error){
  if(value<0){
    return 0, errors.New("Math:Negative numbers can't be sqrt rooted")
  }
  return math.Sqrt(value), nil
}

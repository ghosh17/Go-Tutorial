package main
import "testing"

func TestHelloGo(t *testing.T){
  //function to test hello
  Req:="Hello !World"
  actual:=Hello()
  if(actual!=Req){
    t.Errorf("Test Failed", Req, actual)
  }
}

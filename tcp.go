package main

import "fmt"
import "io"
import "net"
import "time"
import (

	"errors"
)

func main(){
	ln, err := net.Listen("tcp", ":9000")
	if(err!=nil){
		errors.New("Error:")
	}
	defer ln.Close()

	for{
		conn,err:=ln.Accept()
		if(err!=nil){
			errors.New("Error")
		}
		io.WriteString(conn,fmt.Sprint("Hello World", time.Now(), "\n"))
		conn.Close()
	}
}

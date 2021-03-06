package main


import (
	"fmt"
	"net"

	"errors"
	"bufio"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func main(){
	server, err := net.Listen("tcp", ":9000")
	if(err!=nil){
		errors.New("Error:")
	}
	defer server.Close()

	for{
		conn,err:=server.Accept()
		if(err!=nil){
			errors.New("Error")
		}
		go handleConn(conn)
	}
}

package main


import (
	"fmt"
	"net"

	"errors"
	"bufio"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	i:=0
	for scanner.Scan(){
		ln:=scanner.Text()
		fmt.Println(ln)
		if(i==0){
			method:=strings.Fields(ln)[0]
			fmt.Println("METHOD", method)
		}else{

		}
		i++
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

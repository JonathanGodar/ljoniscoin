package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main(){
	StartServer()
}

func StartServer(){
	addr, err := net.ResolveTCPAddr("tcp", ":8923")
	DieIf(err)

	ln, err  := net.ListenTCP("tcp", addr)
	DieIf(err)
	defer ln.Close()


	fmt.Println("Listener started")


	quit := false
	for quit == false {
		fmt.Println("Waiting for a connection")
		conn, err := ln.AcceptTCP()
		if err != nil {
			panic(err)
		}
		go HandleConnection(conn);
	}
}

func HandleConnection(conn *net.TCPConn){
	fmt.Printf("Recieved a connection from %v\n", conn.RemoteAddr())
	io.Copy(os.Stdout, conn)
	fmt.Printf("\nConnection %v closed\n", conn.RemoteAddr())
}


func DieIf(err error){
	if err != nil {
		panic(err)
	}
}

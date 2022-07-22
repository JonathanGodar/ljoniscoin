package main

import (
	"flag"
	"fmt"
	"net"
)

func main(){
	var remoteAddr string;

	flag.StringVar(&remoteAddr, "rAddr", "remoteAddress", "remoteAdress")
	flag.Parse()


	laddr, err := net.ResolveTCPAddr("tcp", ":")
	DieIf(err)

	raddr, err := net.ResolveTCPAddr("tcp",remoteAddr)
	DieIf(err)


	fmt.Println("Connecting")
	conn, err := net.DialTCP("tcp", laddr, raddr);
	DieIf(err)
	defer conn.Close()
	fmt.Println("Connected")

	fmt.Println("Wrinting message")
	conn.Write([]byte("hello world"))
}

func DieIf(err error)  {
	if err != nil {
		panic(err)
	}
}


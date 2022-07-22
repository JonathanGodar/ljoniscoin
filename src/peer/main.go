package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main(){
	var remoteAddr, localAddr string;

	flag.StringVar(&remoteAddr, "rAddr", "INVALID ADDRESS", "remoteAdress")
	flag.StringVar(&localAddr, "lAddr", "", "localAddress")
	flag.Parse()


	laddr, err := net.ResolveTCPAddr("tcp", localAddr)
	DieIf(err)

	raddr, err := net.ResolveTCPAddr("tcp",remoteAddr)
	DieIf(err)


	fmt.Println("Connecting")

	var conn *net.TCPConn

	lastAttemptAt := time.Now()

	minDiff := 200 * time.Millisecond;
	for attemptsRemaining := 50; attemptsRemaining >= 0; attemptsRemaining-- {
		conn, err = net.DialTCP("tcp", laddr, raddr);
		if err == nil {
			break
		}

		diff :=  time.Now().Sub(lastAttemptAt)
		fmt.Printf("Attempts remaining: %v\n", attemptsRemaining)
		if diff < minDiff {
			time.Sleep(minDiff - diff)
		}
		lastAttemptAt = time.Now()
	}
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


package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	// receive client message
	for {
		// make a slice
		buf := make([]byte, 1024)
		fmt.Printf("server waiting message from client=%v\n", conn.RemoteAddr().String())
		// waiting read client message
		n, err := conn.Read(buf)
		if err == io.EOF {
			fmt.Println("client quit successfully!")
			return
		} else if err != nil {
			fmt.Println("server receive client message by conn err=", err)
			return
		}

		fmt.Println("server read client message done, character num=", n)
		// display message received
		fmt.Printf("server receive message=%v\n", string(buf[:n]))
	}
}

func main() {
	fmt.Println("server listening...")
	listen, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("server listen err=", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept() // waiting connection
		if err != nil {
			fmt.Println("server listen accept err=", err)
		} else {
			fmt.Printf("conn=%v, remot client eaddr=%v\n", conn, conn.RemoteAddr().String())
		}
		// run a goroutine server a accepted connect
		go process(conn)
	}
}

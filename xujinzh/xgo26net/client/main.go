package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client connect error")
		return
	} else {
		fmt.Println("client connect successfully")
	}

	defer conn.Close()

	for {
		status, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("bufio read err=", err)
		} else {
			fmt.Println("bufio read success, status=", status)
		}
		// if user input exit
		status = strings.Trim(status, " \r\n")
		if status == "exit" {
			fmt.Println("client quit...")
			break
		}

		// send message to server
		n, err := conn.Write([]byte(status + "\n"))
		if err != nil {
			fmt.Println("client send message to server err=", err)
		} else {
			fmt.Println("client send message to server successfully, character num=", n)
		}
	}
}

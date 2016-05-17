package main

import (
	"net"
	"fmt"
	"os"
)

func main() {
	service := ":5555"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		fmt.Println("resolve tcp error:", err.Error())
		os.Exit(1)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("listen tcp error:", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("listener success")
	}

	conn, err := listener.Accept()
	for {

		fmt.Println("conn :", conn)
		if err != nil {
			fmt.Println("accept err:", err.Error())
			continue
		}

		data := make([]byte, 512)
		conn.Read(data)

		if err != nil {
			fmt.Println("accept error:", err.Error())
		} else {
			fmt.Println("accept string:", string(data))
		}

		//daytime := time.Now().String()
		//conn.Write([]byte(daytime))


	}
	conn.Close()
}

package main

import (
	"fmt"
	"log"
	"net"
)


func startServer() {
	listener, err := net.Listen("tcp", "localhost:6666")
	checkError(err)
	fmt.Println("listenner success")

	for {

		conn, err := listener.Accept()
		checkError(err)

		doServerStuff(conn)
	}
}


func doServerStuff(conn net.Conn) {
	nameInfo := make([]byte, 512)
	_, err := conn.Read(nameInfo)
	checkError(err)

	for {
		buf := make([]byte, 512)
		_, err := conn.Read(buf)
		flag := checkError(err)
		if flag == 0 {
			break
		}
		fmt.Println(string(buf))
	}
}


func checkError(err error) int {
	if err != nil {
		if err.Error() == "EOF" {

			return 0
		}
		log.Fatal("an error!", err.Error())
		return -1
	}
	return 1
}

func main() {

	startServer()
}
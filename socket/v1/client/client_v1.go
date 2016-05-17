package main

import (
	"net"
	"fmt"
	"os"
	"bufio"
)

func main() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":5555")
	if err != nil {
		fmt.Println("create tcp address error:", err.Error())
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)


	//reader := bufio.NewReader(os.Stdin)
	//data, _, _ := reader.ReadLine()
	for {

		reader := bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()

		if err != nil {
			fmt.Println("connect tcp error:", err.Error())
			os.Exit(1)
		}

		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("conn write error:", err.Error())
			os.Exit(1)
		}



	}


}

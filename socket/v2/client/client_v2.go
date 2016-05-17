package main

import (
"bufio"
"fmt"
"log"
"net"
"os"
"strings"
)


func connectServer() {

	conn, err := net.Dial("tcp", "localhost:6666")
	checkError(err)
	fmt.Println("connecct succes")

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("name:")
	name, _ := inputReader.ReadString('\n')

	trimName := strings.Trim(name, "\r\n")
	conn.Write([]byte(trimName + " ok\n "))
	for {

		input, _ := inputReader.ReadString('\n')
		trimInput := strings.Trim(input, "\r\n")

		if trimInput == "quit" {
			fmt.Println("")
			conn.Write([]byte(trimName + "  "))
			return
		}

		_, err = conn.Write([]byte(trimName + " says " + trimInput))
	}
}


func checkError(err error) {
	if err != nil {
		log.Fatal("an error!", err.Error())
	}
}


func main() {

	connectServer()
}
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	/*var ips = []string{
		"172.27.28.233",
		"172.27.27.78",
		"172.27.28.141",
		"172.27.22.11",
		"172.27.22.191",
		"172.27.27.69",
		"172.27.27.202",
		"172.27.255.255",
	  } */

	//server
	fmt.Println("Launching server.. ")
	ln, _ := net.Listen("tcp", ":8079") //listen on all interfaces

	conn, _ := ln.Accept() //accept connection

	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("Message Recieved:", string(message))
	newmessage := strings.ToUpper(message)
	conn.Write([]byte(newmessage + " CLIENT \n"))

	//client
	conn1, _ := net.Dial("tcp", "192.168.0.105:8079")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Text to send :")
	text, _ := reader.ReadString('\n')
	fmt.Fprintf(conn1, text+"\n")

	message1, _ := bufio.NewReader(conn1).ReadString('\n')
	fmt.Printf("Message from Server:" + message1)

}

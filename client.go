package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

func main() {
	server_addr := flag.String("d", "", "server ip address")
	command := flag.String("c", "", "remote command")
	port := flag.Int("p", 80, "server port number")
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprint(*server_addr, ":", *port))
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(conn, "GET /%v HTTP/1.0\r\n\r\n", *command)
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Response:", status)
}

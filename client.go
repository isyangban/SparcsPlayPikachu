package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

var kKeyMap = map[string]string{
	"a":     "key a",
	"b":     "key b",
	"x":     "key x",
	"y":     "key y",
	"up":    "key up",
	"down":  "key down",
	"left":  "key left",
	"right": "key right",
	"nop":   "",
}

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
	if check_command(*command) {
		fmt.Fprintf(conn, "GET /%v HTTP/1.0\r\n\r\n", *command)
	} else {
		return
	}
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Print("Response:", status)
}

func check_command(c string) bool {
	text, ok := kKeyMap[c]
	if ok {
		fmt.Println(text + " sent")
	} else {
		fmt.Println("Error: unknown key " + c)
	}
	return ok
}

package main

import (
	"fmt"
	"net"

	"log"
)

func SendOverNet(host string, port int) chan string {
	addr, _ :=net.ResolveTCPAddr("tcp",
		fmt.Sprintf("%s:%d", host, port))
	conn, _ := net.DialTCP("tcp", nil,addr)
	source := make(chan string)
	go write(source, conn)
	return source;
}

func write(source chan string, sink *net.TCPConn){
	for part := range source{
		sink.Write([]byte(part))
	}
	sink.Close()
}
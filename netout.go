package main

import (
	"fmt"
	"net"
	"errors"
)

func SendOverNet(host string, port int) (chan string, error) {
	addr, _ :=net.ResolveTCPAddr("tcp",
		fmt.Sprintf("%s:%d", host, port))
	conn, err := net.DialTCP("tcp", nil,addr)
	if(err != nil){
		return nil, errors.New(fmt.Sprintf("cant connect to %v:%v", host, port))
	}
	source := make(chan string)
	go write(source, conn)
	return source, nil;
}

func write(source chan string, sink *net.TCPConn){
	for part := range source{
		sink.Write([]byte(part))
	}
	sink.Close()
}
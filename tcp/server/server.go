package main

import (
	"fmt"
	"net"
)

func runner(conn net.Conn) {
	var se = [128]byte{}
	for {
		n, err := conn.Read(se[:])
		if err != nil {
			fmt.Println("read error: ", err)
		}
		fmt.Println(string(se[6:n]), "来自客户端---", string(se[:6]))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:23456")
	if err != nil {
		panic("server start error:" + err.Error())
	}
	fmt.Println("server is listening....")

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic("listen err:" + err.Error())
		}
		go runner(conn)
	}
}

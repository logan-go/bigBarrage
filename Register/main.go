package main

import (
	"bufio"
	"net"

	"github.com/logan-go/bigBarrage/common/connectkeeper"
)

var connectKeeperList []connectkeeper.ConnectKeeper

func init() {
	RegisterConfig = &RegisterConf{}
}

func main() {
	connectKeeperList = make(connectkeeper.ConnectKeeper, 0)
	ln, err := net.Listen("tcp", RegisterConfig.Port)
	for {
		conn, _ := ln.Accept()
		go holdConnects(conn)
	}
}

func holdConnects(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	msg := ""
	for {
		msg = ""
		msg = reader.ReadString(byte(30))
	}
}

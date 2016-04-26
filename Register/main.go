package main

import (
	"net"

	"github.com/logan-go/bigBarrage/common/connectkeeper"
)

var connectKeeperList []connectkeeper.ConnectKeeper

func main() {
	connectKeeperList = make(connectkeeper.ConnectKeeper, 0)
	ln, err := net.Listen("tcp", 2346)
	for {
		conn, _ := ln.Accept()
		go holdConnects(conn)
	}
}

func holdConnects(conn net.Conn) {

}

package main

import (
	"net"

	"github.com/logan-go/bigBarrage/common/gateway"
)

var gatewayList []gateway.Gateway

func main() {
	gatewayList = make(gateway.Gateway, 0)
	ln, err := net.Listen("tcp", 2346)
	for {
		conn, _ := ln.Accept()
		go holdConnects(conn)
	}
}

func holdConnects(conn net.Conn) {

}

package main

import (
	"crypto/md5"
	"fmt"
	"net"
	"os"

	"github.com/logan-go/bigBarrage/ConnectKeeper/runtime"
	"github.com/logan-go/bigBarrage/common/connectkeeper"
)

const (
	DEFAULT_LISTEN_PORT = "2346"
)

func main() {
	ln, err := net.Listen("tcp", ":"+DEFAULT_LISTEN_PORT)
	if err != nil {
		fmt.Sprintln(os.Stderr, "Init Failed:"+err.Error())
		os.Exit(1)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Sprintln(os.Stderr, "Accept Failed:"+err.Error())
		}
		go AcceptConn(conn)
	}
}

func AcceptConn(conn net.Conn) {
	defer conn.Close()
	clientId := md5.Sum([]byte(fmt.Sprintf("%X", conn)))
	bClient := &connectkeeper.BarrageConn{}
	bClient.ClientID = clientId
	bClient.Conn = &conn
	runtime.MainConnList[clientId] = bClient
}

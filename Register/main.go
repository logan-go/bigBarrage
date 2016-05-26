package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/logan-go/bigBarrage/common/communication"
	"github.com/logan-go/bigBarrage/common/connectkeeper"
)

var connectKeeperList []connectkeeper.ConnectKeeper

const (
	DEFAULT_CONFIG_FILE = "register.json"
)

func init() {
	//加载配置
	file, err := os.Open(DEFAULT_CONFIG_FILE)
	if err != nil {
		fmt.Sprintf(os.Stderr, "Load config file failed : "+err.Error())
		os.Exit(1)
	}
	content := make([]byte, 0)
	file.Read(&content)

	RegisterConfig = &RegisterConf{}
	json.Unmarshal(content)

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
		msg = reader.ReadBytes(communication.NORMAL_SEPARATOR)
	}
}

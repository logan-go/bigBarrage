package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/logan-go/bigBarrage/Register/runtime"
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
		fmt.Fprintln(os.Stderr, "Load config file failed : "+err.Error())
		os.Exit(1)
	}
	content := make([]byte, 1024)
	file.Read(content)
	fmt.Println(string(content))

	/**
	fmt.Printf("%v\n", runtime.RegisterConfig)
	json.Unmarshal(content, &runtime.RegisterConfig)
	fmt.Printf("%v\n", runtime.RegisterConfig)
	**/
	runtime.RegisterConfig.BootLogFile = "aaaa"
	runtime.RegisterConfig.ErrorLogFile = "bbb"
	runtime.RegisterConfig.OpenHost = "0.0.0.0"
	runtime.RegisterConfig.Port = 2346
	aaa, err := json.Marshal(runtime.RegisterConfig)
	fmt.Println(string(aaa), err)
}

func main() {
	/**
	connectKeeperList = make([]connectkeeper.ConnectKeeper, 0)
	ln, _ := net.Listen("tcp", runtime.RegisterConfig.OpenHost+fmt.Sprint(runtime.RegisterConfig.Port))
	for {
		conn, _ := ln.Accept()
		go holdConnects(conn)
	}
	**/
}

func holdConnects(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	msg := []byte("")
	for {
		msg = []byte("")
		msg, _ = reader.ReadBytes(communication.NORMAL_SEPARATOR)
		fmt.Println(msg)
	}
}

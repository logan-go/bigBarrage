package main

import (
	"github.com/logan-go/bigBarrage/common/communication"
	"github.com/logan-go/bigBarrage/common/connectkeeper"
)

var RoomList map[communication.RoomID][]*connectkeeper.ConnectKeeper

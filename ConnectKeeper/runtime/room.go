package runtime

import (
	"github.com/logan-go/bigBarrage/common/communication"
	"github.com/logan-go/bigBarrage/common/connectkeeper"
)

//当前ConnectKeeper所有房间列表，以及房间内ClientID列表
var RoomList map[communication.RoomID][]*connectkeeper.BarrageConn

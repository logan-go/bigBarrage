package runtime

import (
	"github.com/logan-go/bigBarrage/common/communication"
	"github.com/logan-go/bigBarrage/common/connectkeeper"
)

//当前ConnectKeeper所有客户端连接（不包括服务间通信）的集合
var MainConnList map[communication.ClientID]*connectkeeper.BarrageConn

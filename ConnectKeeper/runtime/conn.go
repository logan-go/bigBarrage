package runtime

import "github.com/logan-go/bigBarrage/common/communication"

//当前ConnectKeeper所有客户端连接（不包括服务间通信）的集合
var MainConnList map[communication.ClientID]*connectKeeper.BarrageConn

package runtime

import (
	"github.com/logan-go/bigBarrage/Register/config"
	"github.com/logan-go/bigBarrage/common/communication"
)

//房间列表
//记录每个房间分别在那些服务器上
var RoomList map[communication.RoomID]communication.RoomInfomation

//Register配置
var RegisterConfig config.RegisterConf

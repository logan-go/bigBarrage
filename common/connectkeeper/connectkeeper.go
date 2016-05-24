package connectkeeper

import (
	"net"

	"github.com/logan-go/bigBarrage/common/communication"
)

type ConnectKeeper struct {
	Host    string           //服务器地址
	Port    string           //服务器端口
	RoomMap map[string][]int //
}

//ConnectKeeper使用的连接管理程序
type BarrageConn struct {
	Conn     net.Conn               //长连接资源
	ClientID communication.ClientID //链接ID
	RoomID   communication.RoomID   //连接所在的房间ID
}

//发送String类型的消息
func (bc BarrageConn) SendMsg(msg string) (uint64, error) {

}

//发送[]byte类型的消息
func (bc BarrageConn) SendMsg(msg []byte) (uint64, error) {

}

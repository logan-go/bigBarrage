package communication

import "crypto/md5"

//房间ID
type RoomID uint64

//当前Keeper的名字
type KeeperName string

//用户ID
type MemberID uint64

//客户端连接的ID，由连接的内存地址MD5后获得
type ClientID [md5.Size]byte

//明星魅力值
type StarCharm struct {
	Charm    uint64 //魅力值
	StarName string //明星名字
}

package communication

const (
	MESSAGE_CONNECT_KEEPER_REGISTER = iota
	MESSAGE_REPORT_ROOM_INFO
	MESSAGE_REQUEST_ROOM_INFO
	MESSAGE_SEND_MESSAGE
	MESSAGE_HART_BEAT

	SERVER_CONNECT_KEEPER = iota
	SERVER_REGISTER
	SERVER_SENDER
)

//通讯消息体
//所有服务器间通讯的主结构
type Message struct {
	Type int    //信息类型，类型见MESSAGGE_前缀的说明
	Body string //信息内容
}

//房间信息
//用于服务期间房间信息传递和记录
type RoomInfomation struct {
	RoomID         string             //房间ID
	UserCount      uint64             //房间用户所有服务器统计总数
	RoomServerList []ServerInfomation //房间所在服务器列表
}

//消息发送通知体
//主要用于Sender把需要发送消息的信息发给ConnectKeeper时使用
type MessageForSend struct {
	RoomID      string //房间类型
	SendAll     bool   //是否发送给所有人
	SendPercent int    //发送百分比
}

//服务器信息
//用于各个服务想Register传递自己当前状态信息使用
type ServerInfomation struct {
	Host string //地址
	Port string //端口
	Type int    //服务器类型，见SERVER_前缀的描述
}

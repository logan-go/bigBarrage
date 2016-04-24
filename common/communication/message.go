package communication

const (
	MESSAGE_GETWAY_REGISTER = iota
	MESSAGE_REPORT_ROOM_INFO
	MESSAGE_REQUEST_ROOM_INFO
	MESSAGE_SEND_MESSAGE
	MESSAGE_HART_BEAT

	SERVER_GATEWAY
	SERVER_REGISTER
	SERVER_SENDER
)

type Message struct {
	Type int
	Body string
}

type RoomInfomation struct {
	RoomID    string
	UserCount uint64
	RoomIsIn  []ServerInfomation
}

type MessageForSend struct {
	RoomID      string
	SendAll     bool
	SendPercent int
}

type ServerInfomation struct {
	Host string
	Port string
	Type int
}

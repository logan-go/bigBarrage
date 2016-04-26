package connectkeeper

type ConnectKeeper struct {
	Host    string           //服务器地址
	Port    string           //服务器端口
	RoomMap map[string][]int //
}

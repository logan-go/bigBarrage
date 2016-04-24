package gateway

type Gateway struct {
	Host    string
	Port    string
	RoomMap map[string][]int
}

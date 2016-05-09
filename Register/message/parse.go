package message

import (
	"encoding/json"

	"github.com/logan-go/bigBarrage/common/communication"
)

func Parse(bt_msg []byte) (communication.Message, error) {
	var msg communication.Message
	err := json.Unmarshal(bt_msg, &msg)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

package ipmmessage

import (
	"encoding/json"
	"fmt"
)

var MessageTypeMap = map[string]string{
	"200": "First Presentment",
	"700": "Fee Collection",
	"697": "File Header",
	"695": "File Trailer",
}

type IPMMessage struct {
	DE04 string `json:"de04,omitempty"`
	DE24 string `json:"de24,omitempty"`
	DE71 string `json:"de71,omitempty"`
	// ...
	DE48 DE48 `json:"de48,omitempty"`
}

type DE48 struct {
	PDS0105 string `json:"pds0105,omitempty"`
	PDS0137 string `json:"pds0137,omitempty"`
	PDS220  string `json:"pds220,omitempty"`
	PDS1002 string `json:"pds1002,omitempty"`
}

func (message IPMMessage) GetMessageNumber() string {
	return message.DE71
}

func (message IPMMessage) GetMessageType() string {
	messageType := MessageTypeMap[message.DE24]
	return messageType
}

func (message IPMMessage) PrintIPMMessage() {
	marshaledFP, _ := json.Marshal(message)
	fmt.Println(string(marshaledFP))
}

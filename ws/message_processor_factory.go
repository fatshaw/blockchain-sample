package ws

import "github.com/gorilla/websocket"

func NewProcessor(c *websocket.Conn, message *BlockchainWsResponse) MessageProcessor {

	switch message.MessageType {
	case QUERYLATEST:
		return &QueryLatestProcessor{BaseProcessor{c}}
	case QUERYALL:
		return &QueryAllProcessor{BaseProcessor{c}}
	case RESPONSEBLOCKCHAIN:
		return &ResponseProcessor{BaseProcessor{c}}
	case PING:
		return &PongProcessor{BaseProcessor{c}}
	default:
		return &DummyProcessor{}
	}
}

package ws

type MessageProcessor interface {
	OnMessage(message *BlockchainWsResponse)
}

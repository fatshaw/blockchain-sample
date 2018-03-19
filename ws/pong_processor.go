package ws

type PongProcessor struct {
	BaseProcessor
}

func (processor *PongProcessor) OnMessage(message *BlockchainWsResponse) {

	writesocket(processor.c,PONG, nil)

}

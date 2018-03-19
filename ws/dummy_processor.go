package ws

type DummyProcessor struct {
	BaseProcessor
}

func (processor *DummyProcessor) OnMessage(message *BlockchainWsResponse) {

}
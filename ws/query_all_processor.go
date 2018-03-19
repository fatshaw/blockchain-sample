package ws

import (
	. "github.com/fatshaw/blockchain-sample/blockchain"
)

type QueryLatestProcessor struct {
	BaseProcessor
}

func (processor *QueryLatestProcessor) OnMessage(message *BlockchainWsResponse) {
	writesocket(processor.c, RESPONSEBLOCKCHAIN, []Block{BlockchainInstance.LastBlock()})
}

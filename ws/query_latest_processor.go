package ws

import (
	. "github.com/fatshaw/blockchain-sample/blockchain"
)

type QueryAllProcessor struct {
	BaseProcessor
}

func (processor *QueryAllProcessor) OnMessage(message *BlockchainWsResponse) {
	writesocket(processor.c, RESPONSEBLOCKCHAIN, BlockchainInstance.Blocks)
}

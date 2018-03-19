package ws

import (
	"sort"
	. "github.com/fatshaw/blockchain-sample/blockchain"
	. "github.com/fatshaw/blockchain-sample/util"
)

type ResponseProcessor struct {
	BaseProcessor
}

func (processor *ResponseProcessor) OnMessage(message *BlockchainWsResponse) {

	sort.Sort(ByIndex(message.Blocks))
	latestBlockReceived := message.Blocks[len(message.Blocks)-1]
	latestBlockHeld := BlockchainInstance.LastBlock()
	if latestBlockReceived.Index > latestBlockHeld.Index {
		if latestBlockHeld.Hash == latestBlockReceived.PrevHash {
			BlockchainInstance.AppendBlock(latestBlockReceived)
		} else if len(message.Blocks) == 1 {
			processor.queryAllChain()
		} else {
			processor.replaceChain(message)
		}
	}
}

func (processor *ResponseProcessor) replaceChain(message *BlockchainWsResponse) {
	Logger.Info("ReplaceChain")
	BlockchainInstance.ReplaceChain(&Blockchain{message.Blocks})
	Broadcast(RESPONSEBLOCKCHAIN, []Block{BlockchainInstance.LastBlock()})
}

func (processor *ResponseProcessor) queryAllChain() {
	Logger.Info("QueryAll Blockchain")
	Broadcast(QUERYALL, nil)
}

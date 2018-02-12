package blockchain

import (
	"time"
)

// BlockchainInstance A Blockchain Instance
var BlockchainInstance = Blockchain{}

func init() {
	BlockchainInstance.AppendBlock(genesisBlock())
}

type Blockchain struct {
	Blocks []Block
}

func (blockchain *Blockchain) AppendBlock(block Block) {
	blockchain.Blocks = append(blockchain.Blocks, block)
}

func (blockchain *Blockchain) GenerateBlock(BPM int) (Block, error) {
	var newBlock Block
	newBlock.Index = blockchain.LastBlock().Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.BPM = BPM
	newBlock.PrevHash = blockchain.LastBlock().Hash
	newBlock.Hash = newBlock.CalculateHash()
	return newBlock, nil
}

func (blockchain *Blockchain) LastBlock() Block {
	return blockchain.Blocks[len(blockchain.Blocks)-1]
}

func (blockchain *Blockchain) ReplaceChain(newBlockchain *Blockchain) {
	if len(blockchain.Blocks) < len(newBlockchain.Blocks) {
		blockchain.Blocks = newBlockchain.Blocks
	}
}

func genesisBlock() Block {
	return Block{0, "0", 0, "f78b037f6d1ecfc5a00bc7d96858bdc7af9ac8dbf95fdd5736d0f950ab279b9e", ""}
}

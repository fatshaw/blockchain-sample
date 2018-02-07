package blockchain

import (
	"time"
)

// BlockchainInstance A Blockchain Instance
var BlockchainInstance = Blockchain{}

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

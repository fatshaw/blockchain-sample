package main

import "time"

type Blockchain struct {
	Blocks []Block
}

func (blockchain *Blockchain) AppendBlock(block Block) {
	blockchain.Blocks = append(blockchain.Blocks, block)
}

func (blockchain *Blockchain) GenerateBlock(BPM int) (Block, error) {
	var newBlock Block
	newBlock.Index = blockchain.lastBlock().Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.BPM = BPM
	newBlock.PrevHash = blockchain.lastBlock().Hash
	newBlock.Hash = newBlock.CalculateHash()
	return newBlock, nil
}

func (blockchain *Blockchain) lastBlock() Block {
	return blockchain.Blocks[len(blockchain.Blocks)-1]
}

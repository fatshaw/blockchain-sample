package blockchain

import "fmt"

type BigBlock struct {
	Block
	big int
}

func (block *BigBlock) Print() string {
	return fmt.Sprintf("index=%d,timestamp=%s,bpm=%d,hash=%s,prehash=%s,big=%d", block.Index, block.Timestamp, block.BPM, block.Hash, block.PrevHash, block.big)
}

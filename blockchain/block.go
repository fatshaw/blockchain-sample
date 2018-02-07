package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
)

type Block struct {
	Index     int
	Timestamp string
	BPM       int
	Hash      string
	PrevHash  string
}

func (block *Block) CalculateHash() string {
	record := string(block.Index) + block.Timestamp + string(block.BPM) + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (block *Block) IsValidBlock(oldBlock Block) bool {
	if block.Index != oldBlock.Index+1 {
		return false
	}

	if block.PrevHash != oldBlock.Hash {
		return false
	}

	if block.Hash != block.CalculateHash() {
		return false
	}

	return true
}

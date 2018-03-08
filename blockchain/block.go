package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"math/rand"
	"time"
	. "github.com/fatshaw/blockchain-sample/util"
	"strconv"
)

type Block struct {
	Index      int
	Timestamp  string
	BPM        int
	Hash       string
	PrevHash   string
	Difficulty int
	Nonce      int
}

func (block *Block) String() string {
	return fmt.Sprintf("index=%d,timestamp=%s,bpm=%d,hash=%s,prehash=%s,nonce=%d,difficulty=%d", block.Index,
		block.Timestamp, block.BPM, block.Hash, block.PrevHash, block.Nonce, block.Difficulty)
}

func (block *Block) CalculateHash() string {
	record := strconv.Itoa(block.Index) + block.Timestamp + strconv.Itoa(block.BPM) + block.PrevHash + strconv.Itoa(block.Nonce)
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

func (block *Block) TryGenerateHash() {
	for {
		block.Nonce = rand.Int()
		block.Hash = block.CalculateHash()
		if strings.HasPrefix(block.Hash, strings.Repeat("0", difficulty)) {
			Logger.Info("well done, we got a block = ", block)
			break
		}
		Logger.Infof("block.Hash=%s,nonce=%d,sleep 50ms to retry work ...", block.Hash, block.Nonce)
		time.Sleep(time.Millisecond * 50)
	}
}

type ByIndex []Block

func (a ByIndex) Len() int           { return len(a) }
func (a ByIndex) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByIndex) Less(i, j int) bool { return a[i].Index < a[j].Index }

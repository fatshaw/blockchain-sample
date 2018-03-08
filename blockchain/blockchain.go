package blockchain

import (
	"time"
	"sync"
	. "github.com/fatshaw/blockchain-sample/util"
)

// BlockchainInstance A Blockchain Instance
var BlockchainInstance = Blockchain{}

var difficulty = 1
var numOfAvgBlockExpectedWithinTimespan = 1
var timeIntervalOfChangeDifficulty = time.Second * 30
var numOfBlockGeneratedWithinTimespan = 0

var mutex = sync.Mutex{}

func init() {
	BlockchainInstance.AppendBlock(genesisBlock())
	go func() {
		for {

			mutex.Lock()
			difficulty = numOfBlockGeneratedWithinTimespan / numOfAvgBlockExpectedWithinTimespan
			if difficulty == 0 {
				difficulty = 1
			}

			Logger.Infof("timeIntervalOfChangeDifficulty=%d,numOfBlockGeneratedWithinTimespan=%d,numOfAvgBlockExpectedWithinTimespan=%d, "+
				"Change difficulty to %d", timeIntervalOfChangeDifficulty, numOfBlockGeneratedWithinTimespan,
				numOfAvgBlockExpectedWithinTimespan, difficulty)

			numOfBlockGeneratedWithinTimespan = 0
			mutex.Unlock()

			time.Sleep(timeIntervalOfChangeDifficulty)
		}
	}()
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
	newBlock.Difficulty = difficulty
	newBlock.TryGenerateHash()

	mutex.Lock()
	numOfBlockGeneratedWithinTimespan++
	defer mutex.Unlock()

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
	return Block{0, "0", 0, "0a1e2340446d3e7dd298d5ed01c102811d4f7da34255f287b242d0eb7471e352", "", difficulty, 0}
}

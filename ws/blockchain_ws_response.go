package ws

import (
	. "github.com/fatshaw/blockchain-sample/blockchain"
)

const (
	QUERYLATEST        = 0
	QUERYALL           = 1
	RESPONSEBLOCKCHAIN = 2
	PING               = 3
	PONG               = 4
)

type BlockchainWsResponse struct {
	MessageType int
	Blocks      [] Block
}

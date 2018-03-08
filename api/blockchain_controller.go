package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	. "github.com/fatshaw/blockchain-sample/blockchain"
	. "github.com/fatshaw/blockchain-sample/ws"
	. "github.com/fatshaw/blockchain-sample/util"
)

func GetBlockChain(c *gin.Context) {
	c.JSON(http.StatusOK, BlockchainInstance)
}

func SaveBlockChain(c *gin.Context) {
	var addBlockCommand AddBlockCommand
	if err := c.ShouldBindWith(&addBlockCommand, binding.JSON); err == nil {
		newBlock, err := BlockchainInstance.GenerateBlock(addBlockCommand.BPM)
		if err != nil {
			c.JSON(http.StatusInternalServerError, addBlockCommand)
			return
		}

		if newBlock.IsValidBlock(BlockchainInstance.LastBlock()) {
			BlockchainInstance.AppendBlock(newBlock)
			Broadcast(RESPONSEBLOCKCHAIN, []Block{BlockchainInstance.LastBlock()})
			Logger.Info(ToJson(BlockchainInstance))
		}

		c.JSON(http.StatusCreated, newBlock)

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func GetPeer(c *gin.Context) {
	address := make([]string, len(Peers))
	for _, c := range Peers {
		address = append(address, c.RemoteAddr().String())
	}
	c.JSON(http.StatusOK, address)
}

func AddPeer(c *gin.Context) {
	var addPeerCommand AddPeerCommand
	if err := c.ShouldBindWith(&addPeerCommand, binding.JSON); err == nil {
		Connect(&addPeerCommand.Addr)
	}
	c.JSON(http.StatusCreated, addPeerCommand)
}

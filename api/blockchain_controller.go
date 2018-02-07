package api

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/fatshaw/blockchain-sample/blockchain"
)

func GetBlockChain(c *gin.Context) {
	c.JSON(http.StatusOK, blockchain.BlockchainInstance)
}

func SaveBlockChain(c *gin.Context) {
	var message Message
	if err := c.ShouldBindWith(&message, binding.JSON); err == nil {
		newBlock, err := blockchain.BlockchainInstance.GenerateBlock(message.BPM)
		if err != nil {
			c.JSON(http.StatusInternalServerError, message)
			return
		}

		if newBlock.IsValidBlock(blockchain.BlockchainInstance.LastBlock()) {
			blockchain.BlockchainInstance.AppendBlock(newBlock)
			spew.Dump(blockchain.BlockchainInstance)
		}

		c.JSON(http.StatusCreated, newBlock)

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

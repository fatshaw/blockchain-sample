package main

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetBlockChain(c *gin.Context) {
	c.JSON(http.StatusOK, BlockchainInstance)
}

func SaveBlockChain(c *gin.Context) {
	var message Message
	if err := c.ShouldBindWith(&message, binding.JSON); err == nil {
		newBlock, err := BlockchainInstance.GenerateBlock(message.BPM)
		if err != nil {
			c.JSON(http.StatusInternalServerError, message)
			return
		}

		if newBlock.IsValidBlock(BlockchainInstance.lastBlock()) {
			BlockchainInstance.AppendBlock(newBlock)
			spew.Dump(BlockchainInstance)
		}

		c.JSON(http.StatusCreated, newBlock)

	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

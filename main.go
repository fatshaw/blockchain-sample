package main

import (
	"net/http"
	"time"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/fatshaw/blockchain-sample/blockchain"
	"github.com/fatshaw/blockchain-sample/api"
)

func run() {
	router := gin.Default()
	router.GET("/ping", api.Hello)
	router.GET("/blockchains", api.GetBlockChain)
	router.POST("/blockchains", api.SaveBlockChain)

	http.ListenAndServe(":8080", router)
}

func main() {

	go func() {
		blockchain.BlockchainInstance.AppendBlock(blockchain.Block{0, time.Now().String(), 0, "", ""})
		spew.Dump(blockchain.BlockchainInstance)
	}()

	run()
}

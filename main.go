package main

import (
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

var BlockchainInstance Blockchain = Blockchain{}

func run() {
	router := gin.Default()
	router.GET("/ping", hello)
	router.GET("/blockchains", GetBlockChain)
	router.POST("/blockchains", SaveBlockChain)

	http.ListenAndServe(":8080", router)
}

func main() {

	go func() {
		BlockchainInstance.AppendBlock(Block{0, time.Now().String(), 0, "", ""})
		spew.Dump(BlockchainInstance)
	}()

	run()
}

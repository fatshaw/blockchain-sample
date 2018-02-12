package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	. "github.com/fatshaw/blockchain-sample/api"
	. "github.com/fatshaw/blockchain-sample/ws"
	"os"
	"fmt"
	"time"
	. "github.com/fatshaw/blockchain-sample/util"
)


const DEFAULTHTTPPORT = "8080"
const DEFAULTP2PPORT = "6001"

func run() {
	router := gin.Default()
	router.GET("/ping", Hello)
	router.GET("/blockchain/getblocks", GetBlockChain)
	router.POST("/blockchain/mineblock", SaveBlockChain)
	router.GET("/blockchain/getpeers", GetPeer)
	router.POST("/blockchain/addpeer", AddPeer)

	router.Use(gin.LoggerWithWriter(Logger.Writer()))

	http.HandleFunc("/ws/blockchain", BlockchainWsHandler)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         fmt.Sprintf(":%s", getEnvValueWithDefault("P2P_PORT", DEFAULTP2PPORT)),
	}

	go srv.ListenAndServe()
	http.ListenAndServe(fmt.Sprintf(":%s", getEnvValueWithDefault("HTTP_PORT", DEFAULTHTTPPORT)), router)

}

func getEnvValueWithDefault(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		value = defaultValue
	}
	return value
}

func main() {
	run()
}

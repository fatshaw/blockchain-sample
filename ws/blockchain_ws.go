package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
	"encoding/json"
	"sort"
	"net/url"
	"flag"
	"time"
	. "github.com/fatshaw/blockchain-sample/util"
	. "github.com/fatshaw/blockchain-sample/blockchain"
)

var upgrader = websocket.Upgrader{}

var Peers []*websocket.Conn

func init() {
	go func() {
		for {
			Broadcast(PING, nil)
			time.Sleep(5 * time.Second)
		}
	}()
}

func Connect(addr *string) {
	u := url.URL{Scheme: "ws", Host: *flag.String("addr", *addr, "http service address"), Path: "/ws/blockchain"}
	Logger.Info("connecting to", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		Logger.Fatal("dial:", err)
	}

	onConnection(c)
}

func BlockchainWsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		Logger.Info("upgrade:", err)
		return
	}

	Logger.Info("got new connection: ", c.RemoteAddr().String())

	onConnection(c)

}

func onConnection(c *websocket.Conn) {
	Peers = append(Peers, c)
	writesocket(c, QUERYLATEST, nil)
	go onMessage(c)
}

func onMessage(c *websocket.Conn) {
	for {
		_, data, err := c.ReadMessage()
		if err != nil {
			Logger.Info("read from ws error:", err)
			break
		}

		Logger.Info("got new message: ", string(data))



		var message BlockchainWsResponse
		json.Unmarshal(data, &message)

		switch message.MessageType {
		case QUERYLATEST:
			writesocket(c, RESPONSEBLOCKCHAIN, []Block{BlockchainInstance.LastBlock()})
		case QUERYALL:
			writesocket(c, RESPONSEBLOCKCHAIN, BlockchainInstance.Blocks)
		case RESPONSEBLOCKCHAIN:
			handleResponse(&message)
		case PING:
			writesocket(c,PONG,nil)
		}

	}
}


func Broadcast(messageType int, blocks []Block) {
	for _, c := range Peers {
		writesocket(c, messageType, blocks)
	}
}

func handleResponse(message *BlockchainWsResponse) {
	sort.Sort(ByIndex(message.Blocks))
	latestBlockReceived := message.Blocks[len(message.Blocks)-1]
	latestBlockHeld := BlockchainInstance.LastBlock()
	if latestBlockReceived.Index > latestBlockHeld.Index {
		if latestBlockHeld.Hash == latestBlockReceived.PrevHash {
			appendBlock(latestBlockReceived)
		} else if len(message.Blocks) == 1 {
			queryAllChain()
		} else {
			replaceChain(message)
		}
	}
}

func replaceChain(message *BlockchainWsResponse) {
	Logger.Infof("ReplaceChain")
	BlockchainInstance.ReplaceChain(&Blockchain{message.Blocks})
	Broadcast(RESPONSEBLOCKCHAIN, []Block{BlockchainInstance.LastBlock()})
}

func appendBlock(latestBlockReceived Block) {
	Logger.Info("Append block : %v\n", latestBlockReceived)
	BlockchainInstance.AppendBlock(latestBlockReceived)
}

func queryAllChain() {
	Logger.Info("QueryAll Blockchain")
	Broadcast(QUERYALL, nil)
}

func writesocket(c *websocket.Conn, messageType int, blocks []Block) {
	response := BlockchainWsResponse{messageType, blocks}
	s, _ := json.Marshal(&response)
	err := c.WriteMessage(websocket.TextMessage, s)
	if err != nil {
		Logger.Warn("write message error:", err)
	}
}

package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
	"encoding/json"
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
	Logger.Info("connecting to:", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		Logger.Warning("dial:", err)
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
		NewProcessor(c, &message).OnMessage(&message)

	}
}

func Broadcast(messageType int, blocks []Block) {
	for _, c := range Peers {
		writesocket(c, messageType, blocks)
	}
}

func writesocket(c *websocket.Conn, messageType int, blocks []Block) {
	response := BlockchainWsResponse{messageType, blocks}
	s, _ := json.Marshal(&response)
	err := c.WriteMessage(websocket.TextMessage, s)
	if err != nil {
		Logger.Warn("write message error:", err)
	}
}

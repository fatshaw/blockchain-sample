package ws

import (
	. "github.com/gorilla/websocket"
)

type BaseProcessor struct {
	c *Conn
}

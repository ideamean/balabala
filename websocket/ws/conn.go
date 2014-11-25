package ws

import (
	"code.google.com/p/go.net/websocket"
)

type Connection struct {
	ws   *websocket.Conn
	send chan string
}

func NewConnection(ws *websocket.Conn) *Connection {
	return &Connection{send: make(chan string, 256), ws: ws}
}

// Read loop for each connection reads a message and broadcasts it
func (c *Connection) Reader(s Server) {
	var reply string
	for {
		if err := websocket.Message.Receive(c.ws, &reply); err != nil {
			break
		}
		//s.broadcast <- reply
		//reply client
		if err := websocket.Message.Send(c.ws, reply); err != nil {
			c.ws.Close()
		}
	}
	c.ws.Close()
}

// Write loop for each connection writes whatever comes across the send channel
func (c *Connection) Writer() {
	for message := range c.send {
		if err := websocket.Message.Send(c.ws, message); err != nil {
			break
		}
	}
	c.ws.Close()
}

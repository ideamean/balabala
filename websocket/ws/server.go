package ws

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
)

type Server struct {
	onlineUserNum int
	connections   map[*Connection]bool
	broadcast     chan string
	register      chan *Connection
	unregister    chan *Connection
}

func NewServer() Server {
	return Server{
		onlineUserNum: 0,
		broadcast:     make(chan string),
		register:      make(chan *Connection),
		unregister:    make(chan *Connection),
		connections:   make(map[*Connection]bool),
	}
}

func (s *Server) GetOnlineUserNum() int {
	return s.onlineUserNum
}

// Adds a connection to the connection map
func (s *Server) Register(c *Connection) {
	s.register <- c
}

func (s *Server) Unregister(c *Connection) {
	s.unregister <- c
}

func (s *Server) AddBroadcast(msg string) {
	s.broadcast <- msg
}

func (s *Server) Run() {
	for {
		select {
		// Adds a connection
		case c := <-s.register:
			fmt.Println("Notice: Websocket Connected")
			s.onlineUserNum += 1
			s.connections[c] = true
		case c := <-s.unregister:
			fmt.Println("Notice: Websocket disConnected")
			s.onlineUserNum -= 1
			delete(s.connections, c)
			close(c.send)
		case msg := <-s.broadcast:
			fmt.Printf("Broadcasting: %s\n", msg)
			for c := range s.connections {
				if err := websocket.Message.Send(c.ws, msg); err != nil {
					delete(s.connections, c)
					go c.ws.Close()
				}
			}
		}
	}
}

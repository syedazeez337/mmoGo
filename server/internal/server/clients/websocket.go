package clients

import (
	"log"

	"github.com/gorilla/websocket"

	"github.com/syedazeez337/mmoGo/server/internal/server"
	"github.com/syedazeez337/mmoGo/server/pkg/packets"
)

type WebSocketClient struct {
	id       uint64
	conn     *websocket.Conn
	hub      *server.Hub
	sendChan chan *packets.Packet
	logger   *log.Logger
}

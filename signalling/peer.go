package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type Peer struct {
	PeerId   string
	Role     string
	RAM      int
	CPU      int
	Conn     *websocket.Conn
	LastBeat time.Time
}

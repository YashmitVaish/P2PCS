package main

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Server struct {
	mu    sync.Mutex
	peers map[string]*Peer
}

func NewServer() *Server {
	return &Server{
		peers: make(map[string]*Peer),
	}
}

func (s *Server) handleConn(conn *websocket.Conn) {
	defer conn.Close()

	log.Println("heyyy")
}

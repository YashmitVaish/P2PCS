package main

import (
	"sync"
	"time"

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

func (s *Server) addPeer(p *Peer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.peers[p.PeerId] = p
}

func (s *Server) removePeer(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.peers, id)
}

func (s *Server) handleListPeers(conn *websocket.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var peers []PeerInfo
	for _, p := range s.peers {
		if p.Role == "executor" {
			peers = append(peers, PeerInfo{
				PeerID: p.PeerId,
				RAM:    p.RAM,
				CPU:    p.CPU,
			})
		}
	}

	conn.WriteJSON(Message{
		Type:  "peers",
		Peers: peers,
	})
}

func (s *Server) forwardMsg(msg Message) {
	s.mu.Lock()
	target, ok := s.peers[msg.To]
	s.mu.Unlock()

	if !ok {
		return
	}
	target.Conn.WriteJSON(msg)
}

func (s *Server) RegisterBeat(msg Message) {
	s.mu.Lock()
	peer, ok := s.peers[msg.PeerID]
	s.mu.Unlock()

	if !ok {
		return
	}

	peer.LastBeat = time.Now()
}

func (s *Server) Cleanup() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			now := time.Now()

			s.mu.Lock()
			for id, p := range s.peers {
				if now.Sub(p.LastBeat) > 30*time.Second {
					p.Conn.Close()
					delete(s.peers, id)
				}
			}
			s.mu.Unlock()
		}
	}()
}

func (s *Server) handleConn(conn *websocket.Conn) {
	defer conn.Close()

	var peer *Peer

	for {
		var msg Message

		if err := conn.ReadJSON(&msg); err != nil {
			if peer != nil {
				s.removePeer(peer.PeerId)
			}
			return
		}

		switch msg.Type {
		case "register":
			peer = &Peer{
				PeerId:   msg.PeerID,
				Role:     msg.Role,
				RAM:      msg.RAM,
				CPU:      msg.CPU,
				Conn:     conn,
				LastBeat: time.Now(),
			}

			s.addPeer(peer)

		case "list_peers":
			s.handleListPeers(conn)

		case "signal":
			s.forwardMsg(msg)
		}

	}

}

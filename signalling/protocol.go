package main

type Message struct {
	Type string `json:"type"`

	PeerID string `json:"peer_id,omitempty"`
	Role   string `json:"role,omitempty"`
	RAM    int    `json:"ram,omitempty"`
	CPU    int    `json:"cpu,omitempty"`

	Peers []PeerInfo `json:"peers,omitempty"`

	To   string      `json:"to,omitempty"`
	From string      `json:"from,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type PeerInfo struct {
	PeerID string `json:"peer_id"`
	RAM    int    `json:"ram"`
	CPU    int    `json:"cpu"`
}

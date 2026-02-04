package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	server := NewServer()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("upgrade error:", err)
			return
		}
		go server.handleConn(conn)
	})

	log.Println("signaling server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

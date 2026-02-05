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

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	body, err := io.ReadAll(r.Body)
	// 	if err != nil {
	// 		http.Error(w, "not able to read", 403)
	// 		return
	// 	}

	// 	w.WriteHeader(200)
	// 	w.Write([]byte(body))
	// })

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

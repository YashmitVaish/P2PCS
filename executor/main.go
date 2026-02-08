// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func main() {
// 	code := `
// import time
// print("hello from executor")
// for i in range(60):
// 	print(i)
// 	time.sleep(1)
// 	`
// 	parent := context.Background()

// 	ctx, cancel := context.WithTimeout(parent, 60*time.Second)
// 	defer cancel()

// 	result, err := RunCode(ctx, code)
// 	if err != nil {
// 		fmt.Println("execution error:", err)
// 		return
// 	}

// 	fmt.Println("Programme Finished with code: ", result.ExitCode)
// }

package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(
		"ws://localhost:8080/ws", nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	peerID := "executor-1"

	conn.WriteJSON(map[string]interface{}{
		"type":    "register",
		"peer_id": peerID,
		"role":    "executor",
		"ram":     32768,
		"cpu":     8,
	})

	log.Println("registered as executor", peerID)

	ticker := time.NewTicker(10 * time.Second)
	for range ticker.C {
		conn.WriteJSON(map[string]string{
			"type":    "ping",
			"peer_id": peerID,
		})
		log.Println("ping sent")
	}
}

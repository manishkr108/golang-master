// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{}
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	userID := r.URL.Query().Get("user") // ?user=123

	// Subscribe to typing events for this user
	ctx := context.Background()
	sub := rdb.Subscribe(ctx, "typing:"+userID)

	// Listen Redis in background
	go func() {
		for msg := range sub.Channel() {
			conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		}
	}()

	// Read from client → publish typing event
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			return
		}
		// Example: {"event":"typing","to":"456"}
		targetID := "456"
		rdb.Publish(ctx, "typing:"+targetID, string(message))
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

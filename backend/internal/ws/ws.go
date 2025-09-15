package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// converts normal connection to websocket connection
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allows all origins, adjust for production use
	}, // In production, we could check the origin here
}

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	clients map[*websocket.Conn]bool // connected clients
	mu      sync.Mutex               // protects the clients map
	//Mutex means mutual exclusion, which is a concurrency primitive used to protect
	//shared resources from being accessed simultaneously by multiple goroutines.
	// It prevents race conditions and ensures data integrity.
}

var hub = Hub{
	clients: make(map[*websocket.Conn]bool), // initialize the clients map
}

// HandleWS New connection handler
func HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	hub.mu.Lock()
	hub.clients[conn] = true
	hub.mu.Unlock()

	for {
		_, _, err := conn.ReadMessage() // we don't care about the message content
		if err != nil {
			break
		}
	}

	hub.mu.Lock()
	delete(hub.clients, conn)
	hub.mu.Unlock()
}

// Broadcast Function to broadcast messages to all connected clients
func Broadcast(message []byte) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	for client := range hub.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			return
		}
	}
}

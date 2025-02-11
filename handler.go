package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var (
	clients = make(map[*websocket.Conn]bool)

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	mutex     sync.Mutex
	broadcast = make(chan map[*websocket.Conn][]byte)
)

var Chat []string
var ChatType int

func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error in upgrade to websocket")
		return
	}

	defer conn.Close()

	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()
	log.Println("Client connected")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading:", err)
			mutex.Lock()
			delete(clients, conn)
			break
		}
		connMap := make(map[*websocket.Conn][]byte)
		connMap[conn] = msg
		broadcast <- connMap
	}
}

func HandleBroadcast() {
	for {
		connMap := <-broadcast

		for conn := range clients {
			mutex.Lock()
			_, ok := connMap[conn]
			if !ok {
				for _, msg := range connMap {
					if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
						log.Println("Error broadcasting:", err)
						delete(clients, conn)
					}
				}
			}
			mutex.Unlock()
		}
	}
}

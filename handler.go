package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	clients = make(map[*websocket.Conn]bool)

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	mutex     sync.Mutex
	broadcast = make(chan broadcastStruct)
)

type broadcastStruct struct {
	Username string
	Conn     *websocket.Conn
	msg      []byte
}

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
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading:", err)
			mutex.Lock()
			delete(clients, conn)
			log.Println("disconnectiong the client")
			mutex.Unlock()
			break
		}
		htmxMap := make(map[string]interface{})
		if err = json.Unmarshal(data, &htmxMap); err != nil {
			log.Println(err)
			break
		}
		msg := htmxMap["payload"].(string)

		msgStruct := broadcastStruct{
			Conn: conn,
			msg:  []byte(msg),
		}
		broadcast <- msgStruct
	}
}

func HandleBroadcast() {
	for {
		msgStruct := <-broadcast

		for conn := range clients {
			mutex.Lock()
			if msgStruct.Conn != conn {
				if err := conn.WriteMessage(websocket.TextMessage, msgStruct.msg); err != nil {
					log.Println("Error broadcasting:", err)
					delete(clients, conn)
				}
			}
			mutex.Unlock()
		}
	}
}

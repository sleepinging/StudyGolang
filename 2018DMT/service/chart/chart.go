package chart

import (
	"../../tools"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message,10)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var persons = 0

func OnChartConnections(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true
	persons++
	broadmsg(fmt.Sprintf("有人进入房间，当前%d人在线", persons))
	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			//log.Printf("error: %v", err)
			persons--
			broadmsg(fmt.Sprintf("有人离开房间，，当前%d人在线", persons))
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func broadmsg(msg string) {
	m := Message{
		Username: "系统消息",
		Message:  msg,
	}
	fmt.Println(tools.FmtTime(), msg)
	broadcast <- m
}

func OnChartMessages() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				//log.Printf("error: %v", err)
				persons--
				broadmsg(fmt.Sprintf("有人离开房间，当前%d人在线", persons))
				client.Close()
				delete(clients, client)
			}
		}
	}
}

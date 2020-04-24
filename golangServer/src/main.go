package main

// https://scotch.io/bar-talk/build-a-realtime-chat-server-with-go-and-websockets CHECK IT!!
import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Message defines our message object
type Message struct {
	// Email    string `json:"email"`
	// Username string `json:"username"`
	Message string `json:"message"`
}

func main() {
	// Create a simple file server
	// fs := http.FileServer(http.Dir("../public"))
	// http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handleConnections...")
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// setupResponse(&w, r)
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		fmt.Println("the request:")
		// fmt.Println(ws)

		reader(ws)
		// var msg Message
		// Read in a new message as JSON and map it to a Message object
		// err := ws.ReadJSON(&msg)
		// fmt.Println(msg)
		// if err != nil {
		// 	log.Printf("error: %v", err)
		// 	delete(clients, ws)
		// 	break
		// }
		// // Send the newly received message to the broadcast channel
		// broadcast <- msg
	}
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// log.Println(messageType)
		// log.Println((p))
		s := string(p)
		fmt.Println(s)

		if err := conn.WriteMessage(1, []byte("Here is a string....")); err != nil {
			log.Println(err)
			return
		}
	}
}

func handleMessages() {
	fmt.Println("handleMessages")

	// for {
	// 	// Grab the next message from the broadcast channel
	// 	msg := <-broadcast
	// 	// Send it out to every client that is currently connected
	// 	for client := range clients {
	// 		err := client.WriteJSON(msg)
	// 		if err != nil {
	// 			log.Printf("error: %v", err)
	// 			client.Close()
	// 			delete(clients, client)
	// 		}
	// 	}
	// }
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
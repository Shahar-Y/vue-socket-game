package main

// https://scotch.io/bar-talk/build-a-realtime-chat-server-with-go-and-websockets CHECK IT!!
import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var sm sync.Map

// var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Position) // broadcast channel

// Configure the upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Position defines our position object
type Position struct {
	X int64
	Y int64
}

var pos = Position{X: 100, Y: 100}

func main() {
	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go broadcastMessage()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	sm.Store(ws, ws)

	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println("err in readMessage!!")
			conn.Close()
			sm.Delete(conn)
			return
		}

		s := string(p)
		fmt.Println(s)

		switch s {
		case "a":
			pos.X -= 5
			break
		case "d":
			pos.X += 5
			break
		case "w":
			pos.Y -= 5
			break
		case "s":
			pos.Y += 5
			break
		}
		fmt.Println(pos)

		if err := conn.WriteJSON(pos); err != nil {
			log.Println(err)
			conn.Close()
			sm.Delete(conn)
			// delete(clients, conn)
			return
		}
		broadcast <- pos
	}
}

func broadcastMessage() {
	fmt.Println("handleMessages")

	for {
		i := 0
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		sm.Range(func(ki, c interface{}) bool {
			i++
			fmt.Printf("i: %v\n", i)
			client, ok := c.(*websocket.Conn)
			if ok != true {
				fmt.Println("NOT OK!")
			}
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				sm.Delete(ki)
				// delete(clients, client)
			}
			return true
		})
	}
}

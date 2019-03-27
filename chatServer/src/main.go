package main

import (
	"log" // logs errors and informational messages
	"net"
	"net/http" // standard library for creating a simple http server
	"sync"     // library for managing synchronous connections

	// gorilla is a library that handles websockets
	"github.com/gorilla/websocket"
)

// using "map" datatype that has convenient methods to add and remove
var clients = make(map[*websocket.Conn]bool)

// channel to send messages to other goroutines
var broadcast = make(chan Message)

// mutex to synchronize access to shared resources
var mut sync.Mutex

// Configure the upgrader
var upgrader = websocket.Upgrader{}

// Define our message object
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("../public")) // set the directory to host the chat server
	http.Handle("/", fs)                         // this is the default route; allows the user to access index.html

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages; this function is concurrent - see the go keyword
	go handleMessages()

	// Print message to the console showing that the server has started
	log.Println("http server started on :8000")
	// My function to get ip address
	log.Println("Server IP: " + getIP())

	// Start the server on localhost port 8000 and log any errors
	err := http.ListenAndServe(":8000", nil)
	// If there is an error, log it and exit the program
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Change HTTP GET request to a websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	// If this connection fails, end this routine and log the error
	if err != nil {
		log.Fatal(err)
	}
	// This closes the WebSocket connection when this function returns
	defer ws.Close()

	// Register our new client to the map variable - using a mutex to protect this shared variable
	mut.Lock()
	clients[ws] = true
	mut.Unlock()

	// this loop waits for new messages to be written to this websocket
	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		// if we get an error here, assume the client has disconnected, and remove them from the clients map
		if err != nil {
			log.Printf("error: %v", err)
			mut.Lock()
			delete(clients, ws)
			mut.Unlock()
			break
		}
		// Send the newly received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	// infinite loop to listen for messages
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		mut.Lock()
		// range is our advanced for loop in GO
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
		mut.Unlock()
	}
}

// get the server's IP address
func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Oops: " + err.Error() + "\n")
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return (ipnet.IP.String() + "\n")
			}
		}
	}
	return ""
}

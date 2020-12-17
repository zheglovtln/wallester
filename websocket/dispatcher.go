package websockets

import (
	"fmt"
	"log"
	"wallester/repository"
)


// Message wraps the relevant information needed to broadcast a message
type Message struct {
	Channel string      // the channel name to broadcast on
	Data    interface{} // the data to broadcast
}

// Dispatcher maintains the set of active clients and broadcasts messages to the
// clients.
type Dispatcher struct {
	// Broadcase messages to all client.
	broadcast chan *Message

	// Registered clients.
	clients map[string]map[*Client]bool

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	repo repository.CustomerRepository
}

// NewDispatcher creates a new Dispatcher
func NewDispatcher(repo repository.CustomerRepository) *Dispatcher {
	return &Dispatcher{
		broadcast:  make(chan *Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]map[*Client]bool),
		repo: repo,
	}
}

// Broadcast returns the broadcast channel
func (d *Dispatcher) Broadcast() chan *Message {
	// NOTE: To scale this, a message queue must be used to "broadcast" messages
	// to all running webservers. This function would enqueue messages to the
	// the message queue system, while the dispatcher also reads messages from the queue
	// and sends them to the `broadcast` channel
	return d.broadcast
}

func (d *Dispatcher) onConnect(client *Client) {
	log.Println("Client connected: ", client.conn.RemoteAddr())
	channel := client.channel
	customer, _ := d.repo.FindById(channel)
	err := d.repo.LockedLock(&customer)
	if err != nil {
		fmt.Println("Update Failed")
	}
	if _, ok := d.clients[client.channel]; !ok {
		d.clients[channel] = make(map[*Client]bool)
	}
	d.clients[channel][client] = true
	fmt.Printf("%v",d.clients)
}

func (d *Dispatcher) onDisconnect(client *Client) {
	log.Println("Client disconnect: ", client.conn.RemoteAddr())
	channel := client.channel
	customer, _ := d.repo.FindById(client.channel)
	err := d.repo.LockedUnlock(&customer)
	if err != nil {
		fmt.Println("Update Failed")
	}
	if _, ok := d.clients[channel][client]; ok {
		delete(d.clients[channel], client)
		close(client.send)
	}
}


// Run starts the dispatch loop
func (d *Dispatcher) Run() {
	for {
		select {
		case client := <-d.register:
			d.onConnect(client)
		case client := <-d.unregister:
			d.onDisconnect(client)
		case message := <-d.broadcast:
			channel := message.Channel
			if clients, ok := d.clients[channel]; ok {
				for client := range clients {
					select {
					case client.send <- message.Data:
					default:
						close(client.send)
						delete(d.clients[channel], client)
					}
				}
			}
		}
	}
}
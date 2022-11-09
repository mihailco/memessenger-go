package ws

import "github.com/mihailco/memessenger/pkg/service"

type Hub struct {
	// Registered clients.
	clients map[int]*Client

	services *service.Service

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub(s *service.Service) *Hub {
	return &Hub{
		services:   s,
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[int]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client.id] = client
		case client := <-h.unregister:
			delete(h.clients, client.id)
			close(client.send)
		case message := <-h.broadcast:
			//вот тут хандлер
			h.WSHandler(message)
			// for _, client := range h.clients {
			// 	select {
			// 	case client.send <- message:

			// 	default:
			// 		close(client.send)
			// 		delete(h.clients, client.id)
			// 	}
			// }
			// for client := range h.clients {
			// 	select {
			// 	case client.send <- message:
			// 	default:
			// 		close(client.send)
			// 		delete(h.clients, client)
			// 	}
			// }
		}
	}
}

package ws

type Room struct {
	ID      int             `json:"id"`
	Name    string          `json:"name"`
	Clients map[int]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[int]*Room
	Register   chan *Client
	Unregister chan *Client
	Notify     chan *Notification
	Messages   chan *ChatMessage
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[int]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Notify:     make(chan *Notification, 5),
		Messages:   make(chan *ChatMessage, 10),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if _, ok := h.Rooms[client.RoomID]; ok {
				r := h.Rooms[client.RoomID]

				if _, ok := r.Clients[client.ID]; !ok {
					r.Clients[client.ID] = client
				}
			}
		case client := <-h.Unregister:
			if _, ok := h.Rooms[client.RoomID]; ok {
				if _, ok := h.Rooms[client.RoomID].Clients[client.ID]; ok {
					if len(h.Rooms[client.RoomID].Clients) > 0 {
						client.Notifications <- &Notification{
							RoomID:   client.RoomID,
							Username: h.Rooms[client.RoomID].Clients[client.ID].Username,
							Type:     "leave",
							Content:  "User left the lobby",
						}
					}

					delete(h.Rooms[client.RoomID].Clients, client.ID)
					close(client.Answers)
					close(client.Notifications)
					close(client.Messages)
				}
			}
		case notification := <-h.Notify:
			if _, ok := h.Rooms[notification.RoomID]; ok {
				for _, client := range h.Rooms[notification.RoomID].Clients {
					client.Notifications <- notification
				}
			}
		case message := <-h.Messages:
			if _, ok := h.Rooms[message.RoomID]; ok {
				for _, client := range h.Rooms[message.RoomID].Clients {
					client.Messages <- message
				}
			}
		}
	}
}

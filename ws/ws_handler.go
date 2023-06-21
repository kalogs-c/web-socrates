package ws

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{hub}
}

type CreateRoomRequest struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	NumberOfQuestions int    `json:"number_of_questions"`
	Timeout           int    `json:"timeout"`
	Subject           string `json:"subject"`
	Capacity          int    `json:"capacity"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var request CreateRoomRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[request.ID] = &Room{
		ID:                request.ID,
		Name:              request.Name,
		Clients:           make(map[int]*Client),
		Answers:           make(chan *Answer),
		Question:          &Question{},
		NumberOfQuestions: request.NumberOfQuestions,
		Timeout:           time.Duration(request.Timeout),
		Subject:           request.Subject,
		Capacity:          request.Capacity,
		Players:           0,
	}

	c.JSON(http.StatusOK, request)
}

var upgrader websocket.Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	roomID, err := strconv.Atoi(c.Param("room_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	username := c.Query("username")
	clientID, err := strconv.Atoi(c.Query("client_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client := &Client{
		Conn:          conn,
		Answers:       make(chan *Answer, 10),
		Notifications: make(chan *Notification, 10),
		Messages:      make(chan *ChatMessage, 10),
		Question:      make(chan *Question),
		RoomID:        roomID,
		ID:            clientID,
		UserID:        userID,
		Username:      username,
	}

	notification := &Notification{
		RoomID:   roomID,
		Username: username,
		Type:     "join",
		Content:  "A new user joined the lobby",
	}

	h.hub.Register <- client
	h.hub.Notify <- notification

	go client.WriteChatMessage()
	go client.ReadChatMessage(h.hub)
	go client.WriteNotification(h.hub)
	go client.WriteQuestion(h.hub)
	client.ReadAnswer(h.hub)
}

type RoomResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) ListRooms(c *gin.Context) {
	rooms := make([]RoomResponse, 0, len(h.hub.Rooms))

	for _, room := range h.hub.Rooms {
		rooms = append(rooms, RoomResponse{
			ID:   room.ID,
			Name: room.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

type ClientResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) ListClients(c *gin.Context) {
	roomId, err := strconv.Atoi(c.Param("room_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, ok := h.hub.Rooms[roomId]; !ok {
		c.JSON(http.StatusOK, []ClientResponse{})
		return
	}

	clients := make([]ClientResponse, 0, len(h.hub.Rooms[roomId].Clients))

	for _, client := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientResponse{
			ID:       client.ID,
			Username: client.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}

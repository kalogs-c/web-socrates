package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn          *websocket.Conn
	Answers       chan *Answer
	Question      chan *Question
	Notifications chan *Notification
	Messages      chan *ChatMessage
	ID            int    `json:"id"`
	RoomID        int    `json:"room_id"`
	UserID        int    `json:"user_id"`
	Username      string `json:"username"`
}

type Answer struct {
	RoomID      int    `json:"room_id"`
	Username    string `json:"username"`
	Alternative int8   `json:"alternative"`
	QuestionID  int    `json:"question_id"`
}

type Notification struct {
	RoomID   int    `json:"room_id"`
	Username string `json:"username"`
	Type     string `json:"type"`
	Content  string `json:"content"`
}

type ChatMessage struct {
	RoomID   int    `json:"room_id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func (c *Client) WriteChatMessage() {
	defer c.Conn.Close()

	for {
		message, ok := <-c.Messages
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadChatMessage(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf("error: %v", err)
			}
			break
		}

		h.Messages <- &ChatMessage{
			RoomID:   c.RoomID,
			Username: c.Username,
			Message:  string(message),
		}
	}
}

func (c *Client) WriteNotification(h *Hub) {
	defer func() {
		h.Unregister <- c
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Notifications
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) WriteQuestion(h *Hub) {
	defer c.Conn.Close()

	for {
		message, ok := <-c.Question
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadAnswer(h *Hub) {
	defer c.Conn.Close()

	for {
		var answer Answer
		err := c.Conn.ReadJSON(&answer)
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err,
				websocket.CloseGoingAway,
				websocket.CloseAbnormalClosure,
			) {
				log.Printf("error: %v", err)
			}
			break
		}

		h.Answers <- &answer
	}
}

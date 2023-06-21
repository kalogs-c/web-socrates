package ws

import (
	"fmt"
	"time"
)

type Room struct {
	ID                int             `json:"id"`
	Name              string          `json:"name"`
	Answers           chan *Answer    `json:"answers"`
	Question          *Question       `json:"question"`
	NumberOfQuestions int             `json:"number_of_questions"`
	Timeout           time.Duration   `json:"timeout"`
	Subject           string          `json:"subject"`
	Clients           map[int]*Client `json:"clients"`
	Players           int             `json:"players"`
	Capacity          int             `json:"capacity"`
}

type Hub struct {
	Rooms      map[int]*Room
	Register   chan *Client
	Unregister chan *Client
	Notify     chan *Notification
	Messages   chan *ChatMessage
	Answers    chan *Answer
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[int]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Notify:     make(chan *Notification, 5),
		Messages:   make(chan *ChatMessage, 10),
		Answers:    make(chan *Answer, 10),
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
				r.Players++

				if r.Players == r.Capacity {
					fmt.Println("All players have joined, start the game")
					h.Notify <- &Notification{
						RoomID:   client.RoomID,
						Username: "Socrates",
						Type:     "start",
						Content:  "All players have joined, start the game",
					}
					questionNumber := 0
					for i := 0; i < r.NumberOfQuestions; i++ {
						fmt.Println("Question number: ", i)
						playersAnswers := h.StartGame(r, &questionNumber)
						for _, client := range r.Clients {
							isCorrect := playersAnswers[client.Username]
							if isCorrect {
								client.Notifications <- &Notification{
									RoomID:   client.RoomID,
									Username: client.Username,
									Type:     "answer",
									Content:  "Correct answer",
								}
							} else {
								client.Notifications <- &Notification{
									RoomID:   client.RoomID,
									Username: client.Username,
									Type:     "answer",
									Content:  "Wrong answer",
								}
							}
						}
					}
				}
			}
		case client := <-h.Unregister:
			if _, ok := h.Rooms[client.RoomID]; ok {
				if _, ok := h.Rooms[client.RoomID].Clients[client.ID]; ok {
					r := h.Rooms[client.RoomID]
					if len(r.Clients) > 0 {
						client.Notifications <- &Notification{
							RoomID:   client.RoomID,
							Username: h.Rooms[client.RoomID].Clients[client.ID].Username,
							Type:     "leave",
							Content:  "User left the lobby",
						}
					}

					delete(r.Clients, client.ID)
					r.Players--
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
		case answer := <-h.Answers:
			if room, ok := h.Rooms[answer.RoomID]; ok {
				room.Answers <- answer
			}
		}
	}
}

func (h *Hub) StartGame(r *Room, questionNumber *int) map[string]bool {
	r.Question = r.Question.GetRandomQuestion(r.Subject)

	timeoutChan := make(chan struct{})
	go func() {
		time.Sleep(r.Timeout)
		timeoutChan <- struct{}{}
	}()

	for _, client := range r.Clients {
		client.Question <- r.Question
	}

	playersAnswers := make(map[string]bool)
	amountOfAnswers := 0

	for {
		select {
		case answer := <-r.Answers:
			if answer.Alternative == r.Question.CorrectAnswer {
				playersAnswers[answer.Username] = true
			} else {
				playersAnswers[answer.Username] = false
			}
			amountOfAnswers++

			if amountOfAnswers == r.Players {
				for _, client := range r.Clients {
					client.Notifications <- &Notification{
						RoomID:   client.RoomID,
						Username: client.Username,
						Type:     "finish",
						Content:  "All players have answered",
					}
				}
				*questionNumber++
				return playersAnswers
			}
		case <-timeoutChan:
			for _, client := range r.Clients {
				client.Notifications <- &Notification{
					RoomID:   client.RoomID,
					Username: client.Username,
					Type:     "timeout",
					Content:  "The game has timed out",
				}
			}
			*questionNumber++
			return playersAnswers
		}
	}
}

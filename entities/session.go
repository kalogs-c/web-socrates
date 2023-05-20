package entities

import "github.com/google/uuid"

type Session struct {
	id      uuid.UUID
	players map[*Player]bool
	join    chan *Player
	leave   chan *Player
	// broadcast chan *Message
}

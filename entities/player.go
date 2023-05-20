package entities

import "github.com/gorilla/websocket"

type Player struct {
	conn *websocket.Conn
}

func NewPlayer(conn *websocket.Conn) *Player {
	return &Player{
		conn: conn,
	}
}

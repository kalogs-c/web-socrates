package ws

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type GameServer struct {
	upgrader *websocket.Upgrader
	sessions map[uuid.UUID]*Session
}

type Command string

const (
	Join   Command = "join"
	Create Command = "create"
)

type Session struct {
	ID         uuid.UUID       `json:"id"`
	Username   string          `json:"username"`
	Connection *websocket.Conn `json:"connection"`
	Command    Command         `json:"command"`
	IsPrivate  bool            `json:"is_private"`
	InLobby    bool            `json:"in_lobby"`
}

func NewGameServer() *GameServer {
	return &GameServer{
		upgrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		sessions: make(map[uuid.UUID]*Session),
	}
}

func (gs *GameServer) Join(session *Session, conn *websocket.Conn) {
	s := gs.sessions[session.ID]
	if s != nil && s.InLobby {
		s.InLobby = false
		return
	}

	fmt.Println(gs.sessions)
	for _, v := range gs.sessions {
		fmt.Println(v)
		if v.InLobby && !v.IsPrivate {
			session.ID = v.ID
			v.InLobby = false
			return
		}
	}
}

func (gs *GameServer) Create(session *Session, conn *websocket.Conn) {
	gs.sessions[session.ID] = session
}

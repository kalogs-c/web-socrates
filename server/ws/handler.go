package ws

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (gs *GameServer) Handler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	session := Session{Connection: conn, ID: uuid.New(), InLobby: true}
	err = conn.ReadJSON(&session)
	if err != nil {
		fmt.Println("ReadJSON err: ", err.Error())
	}

	switch session.Command {
	case Join:
		go gs.Join(&session, conn)
	case Create:
		go gs.Create(&session, conn)
	default:
		fmt.Println("Unknown command")
		return
	}

	for {
		fmt.Println("loop")
		m := struct {
			Message string `json:"message"`
		}{}
		err := session.Connection.ReadJSON(&m)
		if err != nil {
			fmt.Println("ReadJSON err: ", err.Error())
		}
		fmt.Println(m.Message)

		if err := session.Connection.WriteJSON(&m); err != nil {
			fmt.Println("WriteJSON err: ", err)
		}
		for _, v := range gs.sessions {
			if v.ID == session.ID {
				if err := v.Connection.WriteJSON(&m); err != nil {
					fmt.Println("WriteJSON err: ", err)
				}

				fmt.Println(session)
			}
		}
	}
}

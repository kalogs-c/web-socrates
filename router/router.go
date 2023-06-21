package router

import (
	"github.com/gin-gonic/gin"

	"github.com/kalogs-c/web-socrates/entities/user"
	"github.com/kalogs-c/web-socrates/ws"
)

func InitUserRoutes(r *gin.Engine, userHandler *user.Handler) {
	r.POST("/signup", userHandler.CreateUser)
	r.POST("/login", userHandler.Login)
	r.GET("/logout", userHandler.Logout)
}

func InitWSRoutes(r *gin.Engine, wsHandler *ws.Handler) {
	r.POST("/ws/create-room", wsHandler.CreateRoom)
	r.GET("/ws/join-room/:room_id", wsHandler.JoinRoom)
	r.GET("/ws/list-rooms", wsHandler.ListRooms)
	r.GET("/ws/list-clients/:room_id", wsHandler.ListClients)
}

func Run(r *gin.Engine, address string) error {
	return r.Run(address)
}

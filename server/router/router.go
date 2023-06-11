package router

import (
	"github.com/gin-gonic/gin"

	"github.com/kalogs-c/web-socrates/entities/user"
)

var r *gin.Engine

func InitRoutes(userHandler *user.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUser)
}

func Run(address string) error {
	return r.Run(address)
}

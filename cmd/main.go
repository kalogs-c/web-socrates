package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kalogs-c/web-socrates/connections"
	"github.com/kalogs-c/web-socrates/entities/user"
	"github.com/kalogs-c/web-socrates/router"
	"github.com/kalogs-c/web-socrates/ws"
)

func main() {
	db, err := connections.NewPostgresConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepository := user.NewRepository(db.Q)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	r := gin.Default()
	router.InitUserRoutes(r, userHandler)
	router.InitWSRoutes(r, wsHandler)
	router.Run(r, ":8080")
}

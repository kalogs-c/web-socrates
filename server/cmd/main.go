package main

import (
	"github.com/kalogs-c/web-socrates/connections"
	"github.com/kalogs-c/web-socrates/entities/user"
	"github.com/kalogs-c/web-socrates/router"
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

	router.InitRoutes(userHandler)
	router.Run(":8080")
}

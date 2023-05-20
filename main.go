package main

import (
	"github.com/gin-gonic/gin"

	"github.com/kalogs-c/web-socrates/routes"
)

func main() {
	r := gin.Default()
	routes.SetupWebsocketRoutes(r)
	r.Run()
}

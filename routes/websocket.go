package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/kalogs-c/web-socrates/server"
)

func SetupWebsocketRoutes(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new.html", gin.H{
			"title": "WebSocket",
		})
	})

	// gs := ws.NewGameServer()
	r.GET("/ws", server.ServeWS)
}

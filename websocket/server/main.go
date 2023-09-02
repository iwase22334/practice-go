package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	router := gin.Default()

	router.GET("/:room_id/websocket", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			conn.WriteMessage(websocket.TextMessage, []byte("Hello, Websocket"))
			time.Sleep(time.Second)
		}
	})

	router.Run(":8080")

}

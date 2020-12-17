package websockets

import "github.com/gin-gonic/gin"

type Handler interface {
WebsocketHandler(c *gin.Context)
}

func Attach(router *gin.Engine, h Handler) {

router.GET("/ws/:id",h.WebsocketHandler)
}
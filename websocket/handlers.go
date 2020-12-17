package websockets

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Handlers struct{
	dispatcher *Dispatcher
	serveWs func(*Dispatcher, http.ResponseWriter, *http.Request, string)
}

func NewHandlers(dispatcher *Dispatcher) *Handlers {
	return &Handlers{
		dispatcher: dispatcher,
		serveWs: serveWs,
	}
}

func (h *Handlers) WebsocketHandler(c *gin.Context) {
	id := c.Param("id")
	h.serveWs(h.dispatcher,c.Writer, c.Request, id)
}
var WSUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(dispatcher *Dispatcher, w http.ResponseWriter, r *http.Request, channel string) {
	conn, err := WSUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error: %v", err)
		return
	}
	client := NewClient(dispatcher, conn, channel)
	dispatcher.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WriteDispatch()
	go client.ReadDispatch()
}


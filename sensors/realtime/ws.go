package realtime

import (
	"encoding/json"
	"fmt"

	fbs "github.com/antoniodipinto/ikisocket"
	"github.com/gofiber/fiber"
)

//WsManager instance
var wsm WebsocketManager

// Websockets will initialize the WS service
func Websockets(server *fiber.App) {

	wsm = WebsocketManager{
		clients: make(map[string]string),
		sockets: make(map[string]fbs.Websocket),
	}

	server.Use(func(c *fiber.Ctx) {
		//Make the qs user_id available
		c.Locals("user_id", c.Query("user_id"))
		c.Next()
	})

	fbs.On(fbs.EventConnect, func(e *fbs.EventPayload) {
		userID := fmt.Sprintf("%v", e.Kws.Locals("user_id"))
		wsm.AddSocket(e.Kws)
		wsm.AddSocketAlias(userID, e.Kws.UUID)
	})

	fbs.On(fbs.EventDisconnect, func(e *fbs.EventPayload) {
		userID := e.Kws.Locals("user_id")
		wsm.RemoveSocket(e.Kws.UUID)
		if userID != "" {
			wsm.RemoveSocketAlias(userID.(string))
		}
	})

	fbs.On("close", func(payload *fbs.EventPayload) {
		fmt.Println("ws closed" + payload.SocketAttributes["user_id"])
	})

	//TODO: make endpoint configurable
	server.Get("/ws", fbs.New(func(kws *fbs.Websocket) {
		userID := fmt.Sprintf("%v", kws.Locals("user_id"))
		kws.SetAttribute("user_id", userID)

		kws.OnConnect = func() {}
		kws.OnMessage = func(data []byte) {
			message := MessageObject{}
			json.Unmarshal(data, &message)

			if message.To != "" {
				wsm.EmitTo(message.To, data)
			}
		}
	}))
}

//Broadcast send message to all clients
func Broadcast(message []byte) {
	wsm.Broadcast(message)
}

//EmitTo will send a message to a given socket identified by alias
func EmitTo(alias string, data []byte) {
	wsm.EmitTo(alias, data)
}

package realtime

import (
	"fmt"

	"github.com/antoniodipinto/ikisocket"
)

//WebsocketManager websocket helper
type WebsocketManager struct {
	clients map[string]string
	sockets map[string]ikisocket.Websocket
}

//AddSocket  will track a ws socket instance by UUID
func (w *WebsocketManager) AddSocket(socket ikisocket.Websocket) {
	w.sockets[socket.UUID] = socket
}

//AddSocketAlias will map an alias to a socket uuid
func (w *WebsocketManager) AddSocketAlias(alias string, uuid string) {
	w.clients[alias] = uuid
}

//RemoveSocket will remove a ws socket
func (w *WebsocketManager) RemoveSocket(uuid string) {
	if _, ok := w.sockets[uuid]; ok {
		delete(w.sockets, uuid)

		for alias, socketUUID := range w.clients {
			if socketUUID == uuid {
				RemoveSocketAlias(alias)
			}
		}
	}
}

//RemoveSocketAlias remove a socket alias
func (w *WebsocketManager) RemoveSocketAlias(alias string) {
	if _, ok := w.clients[alias]; ok {
		delete(w.clients, alias)
	}
}

//EmitTo will send a message to a given socket identified by alias
func (w *WebsocketManager) EmitTo(alias string, data []byte) {
	if uuid, ok := w.clients[alias]; ok {
		if socket, ok := w.sockets[uuid]; ok {
			err := socket.EmitTo(uuid, data)
			if err != nil {
				fmt.Printf("message send error %s: %v", uuid, err)
			}
		}
	}
}

//Broadcast will send a message to all sockets
func (w *WebsocketManager) Broadcast(message []byte) {
	for _, socket := range w.sockets {
		socket.EmitTo(socket.UUID, message)
	}
}

//MessageObject data struct for WsMessages
type MessageObject struct {
	Data string `json:"data"`
	From string `json:"from"`
	To   string `json:"to"`
}

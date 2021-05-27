package main

import (
	"net/http"

	"time"

	"github.com/gorilla/websocket"

	"./impl"
)

var (
	upgrade = websocket.Upgrader{
		// Allow cross origin access
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	// Add access router and request handler
	http.HandleFunc("/ws", wsRequestHandler)

	// Start HTTP Server on http://localhost:1024/ws
	http.ListenAndServe("0.0.0.0:1024", nil)
}

func wsRequestHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		conn   *impl.Connection
		data   []byte
		err    error
	)

	// Try first handshake and switch to WebSocket protocol
	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil {
		return
	}

	// Capsule raw WebSocket connection to impl WebSocket connection
	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	// Hold persistent connection and test heartbeats
	go func() {
		for {
			if err := conn.WriteMessage([]byte("pong")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Hold persistent connection and do something
	for {
		// Read data from the client
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}

		// Send data to the client
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	// Close persistent connection when errors occur
	conn.Close()
}

package impl

import (
	"errors"

	"sync"

	"github.com/gorilla/websocket"
)

// Connection is a capsuled WebSocket connection
type Connection struct {
	wsConn    *websocket.Conn // Underlying WebSocket connection
	inChan    chan []byte     // Channel for reading message
	outChan   chan []byte     // Channel for writing message
	closeChan chan byte       // Channel for closing connection
	isClosed  bool
	mutex     sync.Mutex
}

// InitConnection is to initialize the connection and start goroutines
func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConn:    wsConn,
		inChan:    make(chan []byte, 1000),
		outChan:   make(chan []byte, 1000),
		closeChan: make(chan byte, 1),
	}

	// Start read goroutine
	go conn.readLoop()

	// Start write goroutine
	go conn.writeLoop()

	return
}

// ReadMessage is a thread-safe API for reading messages from the connection
func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
		// Once connection is closed, report the error
		err = errors.New("connection is closed when reading message")
	}
	return
}

// WriteMessage is a thread-safe API for writing messages to the connection
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		// Once connection is closed, report the error
		err = errors.New("connection is closed when writing message")
	}
	return
}

// Close is a thread-safe API for closing the connection
func (conn *Connection) Close() {
	// Underlying WebSocket connection can reclose multiple times
	conn.wsConn.Close()

	// Channel for closing connection can only close once
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

// readLoop is to keep reading clients' messages
func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)

	for {
		// Read data from the client
		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case conn.inChan <- data: // If everything is ok, push data to the reading channel
		case <-conn.closeChan: // If the closing channel closed, exit the loop
			goto ERR
		}
	}

ERR:
	// Close persistent connection when errors occur
	conn.Close()
}

// writeLoop is to keep sending clients' messages
func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-conn.outChan: // If everything is ok, pull data from the writing channel
		case <-conn.closeChan: // If the closing channel closed, exit the loop
			goto ERR
		}
		// Send data to the client
		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	// Close persistent connection when errors occur
	conn.Close()
}

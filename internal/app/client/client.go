package goclient

import (
	"log"
	"net/url"
	"sync"
	"github.com/gorilla/websocket"

)

type Connection struct {
	id string
	socket *websocket.Conn
}

func New() *Connection {
	con := new(Connection)
	return con
}

func (connection *Connection) Connect(u url.URL, shutdown chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	var err error

	connection.socket, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer connection.socket.Close()

	done := make(chan struct{})

	//Send id to client
	_, id, err := connection.socket.ReadMessage()
	if err != nil {
		log.Println("read id:", err)
		return
	}
	//Store id per connection
	connection.id = string(id)

	go func() {
		defer close(done)
		for {
			_, message, err := connection.socket.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("(%s) recv: %s", connection.id, message)
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-shutdown:
			log.Println("Shutting down connection")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := connection.socket.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			
			<-done
			
			return
		}
	}
}
package sockets

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var CurrentSocket *websocket.Conn

func SetupSockets() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		websocket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		CurrentSocket = websocket
	})

	http.ListenAndServe(":8090", nil)
}

func SendMessage(message string) error {
	if CurrentSocket == nil {
		return errors.New("No socket connected")
	}

	if err := CurrentSocket.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		return err
	}

	return nil
}

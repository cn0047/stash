package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type MyStruct struct {
	MyData string `json:"myData"`
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Fatalf(
					"got error (IsCloseError = %v), error %#v \n",
					websocket.IsCloseError(err, websocket.CloseGoingAway),
					err,
				)
			}

			// Print the message to the console
			fmt.Printf("got from: %v; type: %d; msg: %s\n", conn.RemoteAddr(), msgType, msg)

			ms := MyStruct{MyData: string(msg)}
			d, err := json.Marshal(ms)
			if err != nil {
				log.Fatal(err)
			}

			// Write message back to browser
			err = conn.WriteMessage(websocket.TextMessage, d)
			if err != nil {
				log.Fatal(err)
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.ListenAndServe(":8080", nil)
}

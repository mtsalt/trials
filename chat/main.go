package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle("/ws", websocket.Handler(ws))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ws(ws *websocket.Conn) {
	log.Println("end point : /ws")
	var message string
	for {
		if err := websocket.Message.Receive(ws, &message); err != nil {
			log.Println("ERROR : Failed to receive, ", err)
		}
		if err := websocket.Message.Send(ws, message); err != nil {
			log.Println("ERROR : Failed to send, ", err)
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		so.Join("locationTracking")

		var result map[string]interface{}

		so.On("event:driverLocation", func(msg string) {
			json.Unmarshal([]byte(msg), &result)
			drivertEvent := fmt.Sprintf("%s%s", "event:driver:", result["driver"].(interface{}))
			server.BroadcastTo("locationTracking", drivertEvent, msg)
		})
		so.On("disconnection", func() {
			log.Println(server.Count())
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

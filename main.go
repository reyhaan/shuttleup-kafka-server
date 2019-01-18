package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
	"github.com/reyhaan/go-socket.io-redis"
)

func main() {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	opts := map[string]string{"host": "localhost", "port": "6379"}

	server.SetAdaptor(redis.Redis(opts))
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		so.On("join", func(room string) {
			so.Join(room)
		})

		var result map[string]interface{}

		so.On("event", func(msg string) {
			json.Unmarshal([]byte(msg), &result)

			event := fmt.Sprintf("%s", result["event"].(interface{}))
			data := fmt.Sprintf("%s", result["data"].(interface{}))
			room := fmt.Sprintf("%s", result["room"].(interface{}))

			server.BroadcastTo(room, event, data)
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

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/s4y/ticker/static"
)

func main() {
	httpAddr := flag.String("http", "127.0.0.1:8080", "Listening address")
	flag.Parse()
	fmt.Printf("http://%s/\n", *httpAddr)

	ln, err := net.Listen("tcp", *httpAddr)
	if err != nil {
		log.Fatal(err)
	}

	upgrader := websocket.Upgrader{}

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
		}

		for {
			var msg struct {
				Type      string `json:"type"`
				StartTime uint64 `json:"startTime"`
			}
			if err = conn.ReadJSON(&msg); err != nil {
				break
			}
			switch msg.Type {
			case "ping":
				conn.WriteJSON(struct {
					Type       string `json:"type"`
					StartTime  uint64 `json:"startTime"`
					ServerTime uint64 `json:"serverTime"`
				}{"pong", msg.StartTime, uint64(time.Now().UnixNano()) / uint64(time.Millisecond)})
			default:
				fmt.Println("unknown message:", msg)
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeContent(w, r, "index.html", static.ModTime, strings.NewReader(static.IndexHtml))
		} else if r.URL.Path == "/favicon.ico" {
			return
		} else {
			http.NotFound(w, r)
		}
	})
	log.Fatal(http.Serve(ln, nil))
}

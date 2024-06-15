/*
  A minimalistic demonstration of:
    * serving static content, embedded into your compiled Go binary
		* serving a JSON based API call (and querying it from your page via fetch())
    * serving websocket-based contents (and querying it from your page via WebSocket())

  (C) Robert Kisteleki
	Licensed under MIT license
*/

package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

//go:embed assets
var embeddedFS embed.FS

// we send these objects as JSON
type apiDataStruct struct {
	IntField    int    `json:"i"`
	StringField string `json:"s"`
}

func main() {
	// "assets" is where static stuff goes to, but it's served with HTTP under /
	serverRoot, err := fs.Sub(embeddedFS, "assets")
	if err != nil {
		log.Fatal(err)
	}

	// define the paths we serve: static, API, WS
	http.Handle("/", http.FileServer(http.FS(serverRoot)))
	http.HandleFunc("/api/json/", sampleApiResponse)
	http.Handle("/api/ws/", sampleWsHandle{upgrader: websocket.Upgrader{}})

	// start serving
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func apiData(i int, s string) ([]byte, error) {
	ret := apiDataStruct{i, s}
	return json.Marshal(ret)
}

// a trivial JSON response
func sampleApiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := apiData(0, "X")
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(response))
}

// WebSocket stuff
type sampleWsHandle struct {
	upgrader websocket.Upgrader
}

// here's where the actual WebSocket handler code is
func (wsh sampleWsHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := wsh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error %s when upgrading connection to websocket", err)
		return
	}

	defer func() {
		log.Println("Closing WebSocket")
		conn.Close()
	}()

	for {
		// uncomment if you want to read a message first
		/*
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Error %s when reading message from client", err)
				return
			}
			log.Printf("Received: %s", string(message))
		*/

		// generate some output
		for i := 1; i <= 10; i++ {
			// send one data object
			data, _ := apiData(i, "abcdefghij"[0:i])
			err = conn.WriteMessage(websocket.TextMessage, []byte(data))
			if err != nil {
				log.Printf("Error sending message: %v", err)
				return
			}

			// simulate waiting for some new data to be generated
			time.Sleep(1 * time.Second)
		}
	}
}

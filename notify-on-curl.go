package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/martinlindhe/notify"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"os"
	"strconv"
)

func setup() (int, string) {
	notifyPort := 8080
	if os.Getenv("PORT") != "" {
		var err error
		notifyPort, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
		}
	}
	key := uuid.NewV4()
	return notifyPort, key.String()
}
func serveRequest(w http.ResponseWriter, r *http.Request, key string) {
	switch r.Method {
	case "POST":
		if r.URL.Path == fmt.Sprintf("/%s", key) {
			notify.Notify("Door Cam", "Door Cam", "Someone's at the door", "")
			fmt.Fprintln(w, "Notified!")
			return
		}
		fmt.Fprintln(w, "Error or incorrect key!")
	default:
		fmt.Fprintln(w, "Method not supported!")
	}
}

func main() {
	notifyPort, key := setup()
	fmt.Printf("Serving on port %d, key is %s\n", notifyPort, key)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveRequest(w, r, key)
	})
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(notifyPort),
		handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

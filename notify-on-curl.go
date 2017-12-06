package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/martinlindhe/notify"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func setup() (int, string) {
	rand.Seed(time.Now().UnixNano())
	jokePort := 8080
	if os.Getenv("PORT") != "" {
		var err error
		jokePort, err = strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
		}
	}
	key := make([]rune, 16)
	for i := range key {
		key[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return jokePort, string(key)
}
func serveRequest(w http.ResponseWriter, r *http.Request, key string) {
	switch r.Method {
	case "POST":
		if r.URL.Path == fmt.Sprintf("/%s", key) {
			notify.Notify("Door Cam", "notice", "Someone's at the door", "")
		}
	default:
		fmt.Fprintln(w, "Method not supported!")
	}
}

func main() {
	jokePort, key := setup()
	fmt.Printf("Serving on port %d, key is %s\n", jokePort, key)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveRequest(w, r, key)
	})
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(jokePort),
		handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)))
}

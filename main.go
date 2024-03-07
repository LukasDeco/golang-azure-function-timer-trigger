package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LukasDeco/golang-azure-function-timer-trigger/timertrigger"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	message := "Pass a name in the query string for a personalized response. \n"
	name := r.URL.Query().Get("name")

	if name != "" {
		message = fmt.Sprintf("Hello, %s!\n", name)
	}

	fmt.Fprint(w, message)
}

func timertriggerHandler(w http.ResponseWriter, r *http.Request) {
	var request map[string]any
	d := json.NewDecoder(r.Body)
	d.Decode(&request)

	fmt.Println("got a timer trigger request")

	go timertrigger.Run(context.Background())

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("hey, timer trigger just fired"))
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/timertrigger", timertriggerHandler)
	log.Printf("About to listen on %s. Go to http://127.0.0.1%s", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

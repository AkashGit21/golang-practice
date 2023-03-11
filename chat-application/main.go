package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Start Real-time Chat Application!")
	setupAPI()

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI() {

	manager := NewManager()

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.serveWS)
}

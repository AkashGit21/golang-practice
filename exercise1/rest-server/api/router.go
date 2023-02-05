package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/ready", isReady).Methods("GET")
	return r
}

func isReady(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	w.WriteHeader(200)
	time.Sleep(1 * time.Second)
	w.Write([]byte("I am ready!"))
	log.Printf("Time taken to perform readiness check is %v nanoseconds.", time.Now().Sub(now).Nanoseconds())
}

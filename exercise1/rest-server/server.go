package rest

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AkashGit21/golang-practice/exercise1/rest-server/api"
)

type serverREST struct {
	*http.Server
}

type ServerOperations interface {
	New() (*serverREST, error)
	Start() error
	Shutdown() error
}

// Returns a new HTTP server
func New() (*serverREST, error) {

	// Load the host and port from env here
	host := "localhost"
	port := 8080

	httpSrv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: api.New(),
	}

	return &serverREST{httpSrv}, nil
}

// Starts the HTTP server and enables it to receive requests
func (srv *serverREST) Start() error {
	return nil
}

// Shutdown the server gracefully
func (srv *serverREST) Shutdown() error {
	// Shutdown the REST server
	if srv != nil {
		log.Println("Shutting httpServer!")
		err := srv.Shutdown()
		if err != nil {
			log.Printf("Error while stopping REST server! %v", err)
			return err
		}
	}
	return nil
}

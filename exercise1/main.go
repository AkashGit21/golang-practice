package main

import (
	"log"

	rest "github.com/AkashGit21/golang-practice/exercise1/rest-server"
)

func main() {
	log.Printf("This marks the beginning of exercise 1")
	log.Printf("Choose one of the following server: \n\t1.REST \n\t2.gRPC")

	choice := 1
	switch choice {
	case 1:
		srv, err := rest.New()
		if err != nil {
			panic(err)
		}

		srv.Start()
	case 2:

	}

}

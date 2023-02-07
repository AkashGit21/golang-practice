package main

import (
	"log"
	"time"

	"github.com/AkashGit21/golang-practice/exercise1/verification/utility"
)

const (
	NUMBER_OF_REQUESTS = 1000
	URI                = "http://localhost:8080/ready"
)

func main() {
	log.Printf("Start running GET requests!")

	worker := make(chan struct{}, 1000)
	go func() {
		for index := 0; index < NUMBER_OF_REQUESTS; index++ {
			worker <- struct{}{}
		}
	}()

	time.Sleep(100 * time.Millisecond)
	now := time.Now()

	for len(worker) > 0 {
		go func() {
			// start := time.Now()
			_, err := utility.Get(URI)
			if err != nil {
				log.Printf("Could not receive the Get response: %v", err)
				return
			}
			<-worker
			// log.Printf("Time taken to handle Get request is %v milliseconds.", time.Now().Sub(start).Milliseconds())
			// wg.Done()
		}()
	}
	// log.Printf("%v requests finished!", index+100)
	// }()
	// wg.Wait()
	log.Printf("Total tme taken to run %v GET requests is %v milliseconds.", NUMBER_OF_REQUESTS, time.Now().Sub(now).Milliseconds())
}

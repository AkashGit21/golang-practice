package utility

import (
	"fmt"
	"log"
	"net/http"
)

func Get(uri string) ([]byte, error) {
	var data []byte
	resp, err := http.Get(uri)
	if err != nil {
		log.Printf("Got the following error while performing Get request: %v", err)
	}

	if resp.StatusCode == 200 {
		_, err := resp.Body.Read(data)
		if err != nil {
			return nil, fmt.Errorf("Unable to read the response of Get request: %v", err)
		}
		return data, nil
	}
	return nil, fmt.Errorf("Got the following status code: %v", resp.StatusCode)
}

func Post(uri string, data string, headers string) (string, error) {

	return "", nil
}

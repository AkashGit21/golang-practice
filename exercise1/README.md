## Exercise

## REST Server

To handle millions of requests per minute, you'll need to design an efficient, scalable, and concurrent HTTP server in Golang. Here's a high-level overview of how you could build such a server:

  -  **Parse incoming requests**: Use the net/http library in Golang to parse incoming HTTP requests and extract relevant information such as URL, headers, and body.

  -  **Routing**: Implement a routing mechanism to handle different URLs and HTTP methods.

  -  **Concurrency**: Use Goroutines and Channels to handle multiple requests concurrently and efficiently. You can limit the number of Goroutines created to avoid overloading the system.

  -  **Caching**: Implement caching to store frequently requested data in memory to reduce the number of database queries.

  -  **Load balancing**: If necessary, distribute the incoming requests across multiple servers using load balancing techniques such as round-robin or least connections.

---

## Web Sockets

To handle millions of WebSocket connections per minute in Golang, you need to design an efficient, scalable, and concurrent WebSocket server. Here's a high-level overview of how you could build such a server:

  -  **Parse incoming WebSocket connections**: Use the gorilla/websocket library to parse incoming WebSocket connections and upgrade the HTTP connection to a WebSocket connection.

  -  **Concurrency**: Use Goroutines and Channels to handle multiple connections concurrently and efficiently. You can limit the number of Goroutines created to avoid overloading the system.

  -  **Message** handling: Implement mechanisms to handle messages sent over WebSocket connections, including broadcast, private messages, and message persistence.

  -  **Load balancing**: If necessary, distribute incoming connections across multiple servers using load balancing techniques such as round-robin or least connections.

Here's a sample implementation of a basic WebSocket server in Golang that can handle multiple connections:

```package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            panic(err)
        }
        defer conn.Close()
        for {
            messageType, p, err := conn.ReadMessage()
            if err != nil {
                break
            }
            if err := conn.WriteMessage(messageType, p); err != nil {
                break
            }
        }
    })
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}
```
This code creates a WebSocket server that listens on port 8080 and echoes back all messages sent over a WebSocket connection. You can extend this code to handle millions of connections per minute based on your requirements.


## Caching using Redis

Redis is a popular in-memory data store that can be used as a cache in Golang. Here's an overview of how you can use Redis as a cache in Golang:

  - **Connect to Redis**: Use a Redis client library for Golang, such as Redigo, to establish a connection to a Redis server.

  - **Set and Get data**: Use the SET and GET commands in Redis to store and retrieve data, respectively.

  - **Expiration**: Set an expiration time for cached data using the EXPIRE command in Redis to ensure that the cache does not become too large and consume too much memory.

Here's an example of how you can cache data in Redis using the Redigo library in Golang:
```
package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// set data
	_, err = conn.Do("SET", "key", "value")
	if err != nil {
		panic(err)
	}

	// get data
	value, err := redis.String(conn.Do("GET", "key"))
	if err != nil {
		panic(err)
	}
	fmt.Println("value:", value)
	// set expiration
	_, err = conn.Do("EXPIRE", "key", 3600)
	if err != nil {
		panic(err)
	}
}
```

This code connects to a Redis server on port 6379, sets a key-value pair, retrieves the value for the key, and sets an expiration time for the key. You can use this code as a starting point for your cache implementation using Redis in Golang.


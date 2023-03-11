package main

import "github.com/gorilla/websocket"

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *manager
}

func NewClient(conn *websocket.Conn, manager *manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
	}
}
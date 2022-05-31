package main

import "github.com/J-Obog/go-chat-server/api"

func main() {
	server := api.NewServer("localhost", 8080)
	server.Run()
}

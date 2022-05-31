package main

import (
	"flag"

	"github.com/J-Obog/go-chat-server/api"
)

const (
	defaultPort = 8080
	defaultHost = "localhost"
)

func main() {
	hostFlag := flag.String("h", defaultHost, "host name")
	portFlag := flag.Uint("p", defaultPort, "port to listen on")

	server := api.NewServer(*hostFlag, uint16(*portFlag))
	server.Run()
}

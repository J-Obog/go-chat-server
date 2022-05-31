package main

import (
	"flag"

	"github.com/J-Obog/go-chat-server/api"
)

const (
	defaultPort  = 8080
	defaultHost  = "localhost"
	defaultDebug = false
)

func main() {
	hostFlag := flag.String("h", defaultHost, "host name")
	portFlag := flag.Uint("p", defaultPort, "port to listen on")
	debugFlag := flag.Bool("d", defaultDebug, "debug mode for logging")
	flag.Parse()

	server := api.NewServer(*hostFlag, uint16(*portFlag), *debugFlag)
	server.Run()
}

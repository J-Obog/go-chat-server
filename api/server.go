package api

import "github.com/gorilla/mux"

type Server struct {
	Host    string
	Port    uint16
	handler mux.Router
}

func NewServer(host string, port uint16) *Server {
	server := &Server{Host: host, Port: port}
	router := mux.NewRouter()
	server.handler = *router
	return server
}

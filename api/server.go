package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/J-Obog/go-chat-server/api/service/message"
	"github.com/gorilla/mux"
)

//a server to handle incoming api requests
//handler is initialized with routes from service routers
type Server struct {
	//host name
	Host string

	//port number
	Port    uint16
	handler *mux.Router
}

func NewServer(host string, port uint16) *Server {
	return &Server{
		Host:    host,
		Port:    port,
		handler: initializedHandler(),
	}
}

//adds service handlers to base routers and returns the newly intialized router
func initializedHandler() *mux.Router {
	router := mux.NewRouter()
	message.InitializeMessageRouter(router.PathPrefix("/messages").Subrouter())
	return router
}

//run server
func (this *Server) Run() {
	addr := fmt.Sprintf("%s:%d", this.Host, this.Port)
	log.Fatal(http.ListenAndServe(addr, this.handler))
	os.Exit(0)
}

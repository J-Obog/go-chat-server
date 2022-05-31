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
	Port uint16

	//basic logging turned on when debug enabled
	Debug bool

	handler *mux.Router
}

func NewServer(host string, port uint16, debug bool) *Server {
	server := &Server{Host: host, Port: port, Debug: debug}
	router := mux.NewRouter()
	router.Use(StandardHeadersMiddleware)

	if debug {
		router.Use(LoggingMiddleware)
	}

	//initialize service routes
	message.InitializeMessageRouter(router.PathPrefix("/messages").Subrouter())

	server.handler = router
	return server
}

//run server
func (this *Server) Run() {
	addr := fmt.Sprintf("%s:%d", this.Host, this.Port)
	if this.Debug {
		log.Println("Server listening at " + addr)
	}

	err := http.ListenAndServe(addr, this.handler)

	if err != nil && this.Debug {
		log.Fatal(err)
	}

	os.Exit(0)
}

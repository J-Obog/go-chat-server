package message

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewMessageRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/message", getAllMessages).Methods("GET")
	router.HandleFunc("/message", createNewMessage).Methods("POST")
	return router
}

//get all messages within a specified time frame
func getAllMessages(w http.ResponseWriter, r *http.Request) {

}

//create a new message
func createNewMessage(w http.ResponseWriter, r *http.Request) {

}

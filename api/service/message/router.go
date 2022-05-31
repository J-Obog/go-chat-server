package message

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeMessageRouter(r *mux.Router) {
	r.HandleFunc("/message", getAllMessages).Methods("GET")
	r.HandleFunc("/message", createNewMessage).Methods("POST")
}

//get all messages within a specified time frame
func getAllMessages(w http.ResponseWriter, r *http.Request) {

}

//create a new message
func createNewMessage(w http.ResponseWriter, r *http.Request) {

}

package message

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeMessageRouter(r *mux.Router) {
	r.StrictSlash(true)
	r.HandleFunc("/", getAllMessages).Methods("GET")
	r.HandleFunc("/", createNewMessage).Methods("POST")
}

//get all messages within a specified time frame
func getAllMessages(w http.ResponseWriter, r *http.Request) {
	log.Println("getting messages")
}

//create a new message
func createNewMessage(w http.ResponseWriter, r *http.Request) {

}

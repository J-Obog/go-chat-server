package message

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/segmentio/ksuid"
)

func InitializeMessageRouter(r *mux.Router) {
	InitMessageStore()
	r.StrictSlash(true)
	r.HandleFunc("/", getAllMessages).Methods("GET")
	r.HandleFunc("/", createNewMessage).Methods("POST")
}

//get all messages within a specified time frame
func getAllMessages(w http.ResponseWriter, r *http.Request) {

}

//create a new message
func createNewMessage(w http.ResponseWriter, r *http.Request) {
	var newMessage Message

	err := json.NewDecoder(r.Body).Decode(&newMessage)

	if err != nil {
		w.WriteHeader(500)
		return
	}

	newMessage.ID = ksuid.New()
	newMessage.SentAt = time.Now()
	messageStore = append(messageStore, newMessage)
}

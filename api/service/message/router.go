package message

import (
	"encoding/json"
	"net/http"
	"strconv"
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
	var timestamp time.Time
	messages := []Message{}
	timestr := r.URL.Query().Get("timestamp")

	//only set timestamp if 'timestamp' parameter is present in url
	if timestr != "" {
		stamp, err := strToUinxStamp(timestr)

		if err != nil {
			makeRequestError("Error parsing timestamp", 400, w)
			return
		}
		timestamp = stamp
	}

	//only retrieve messages sent after timestamp
	for _, message := range messageStore {
		if message.SentAt.After(timestamp) {
			messages = append(messages, message)
		}
	}

	err := json.NewEncoder(w).Encode(messages)

	if err != nil {
		makeRequestError("Internal server error", 500, w)
		return
	}
}

//convert a string to a unix timestamp
func strToUinxStamp(str string) (time.Time, error) {
	num, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return time.Unix(0, 0), err
	}

	return time.Unix(num, 0), nil
}

//create a new message
func createNewMessage(w http.ResponseWriter, r *http.Request) {
	var newMessage Message

	err := json.NewDecoder(r.Body).Decode(&newMessage)

	if err != nil {
		makeRequestError("Internal server error", 500, w)
		return
	}

	newMessage.ID = ksuid.New()
	newMessage.SentAt = time.Now()
	messageStore = append(messageStore, newMessage)
	w.Write([]byte("Message successfully stored"))
}

//writing error message from request
func makeRequestError(errMsg string, code int, w http.ResponseWriter) {
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]string{
		"message": errMsg,
		"status":  strconv.Itoa(code),
	})
}

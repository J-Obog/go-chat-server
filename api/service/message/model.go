package message

import (
	"time"

	"github.com/segmentio/ksuid"
)

type Message struct {
	ID      ksuid.KSUID `json:"id"`
	Content string      `json:"content"`
	SentAt  time.Time   `json:"sent_at"`
}

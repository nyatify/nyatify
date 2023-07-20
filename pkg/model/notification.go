package model

import (
	"encoding/json"
	"time"
)

type Notification struct {
	ID    int
	Title string `json:"title"`
	Body  string `json:"body"`

	Client string `json:"client"`

	CreatedAt time.Time `json:"created_at"` // when notification was created. filled by DB
	By        string    `json:"by"`         // who should receive notification

	Rule Rule `json:"rule"` // when to send notification
}

func (n *Notification) String() string {
	return n.Title + ": " + n.Body
}

func (n *Notification) ToJSON() ([]byte, error) {
	return json.Marshal(n)
}

func (n *Notification) FromJSON(data []byte) error {
	return json.Unmarshal(data, n)
}

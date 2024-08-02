package models

import "time"

type Message struct {
    Sender    string    `json:"sender"`
    Recipient string    `json:"recipient"`
    Content   string    `json:"content"`
    Timestamp time.Time `json:"timestamp"`
}

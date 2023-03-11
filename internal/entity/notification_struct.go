package entity

import (
	"database/sql"
	"errors"
	"time"
)

type Notification struct {
	ID         int            `json:"id"`
	Title      string         `json:"title"`
	Message    string         `json:"message"`
	ReceivedAt string         `json:"received_at"`
	OpenedAt   sql.NullString `json:"opened_at"`
}

func NewNotification(title, message string) (*Notification, error) {
	notification := &Notification{
		ID:         0,
		Title:      title,
		Message:    message,
		ReceivedAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := notification.Validate()
	if err != nil {
		return nil, err
	}
	return notification, nil
}

func (n *Notification) Validate() error {
	if n.Title == "" {
		return errors.New("title can't be empty")
	}

	if n.Message == "" {
		return errors.New("message can't be empty")
	}
	return nil
}

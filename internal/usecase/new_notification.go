package usecase

import (
	"database/sql"
	"go-notification-api/internal/entity"
)

type NotificationInputDTO struct {
	Title   string
	Message string
}

type NotificationOutputDTO struct {
	ID         int
	Title      string
	Message    string
	ReceivedAt string
	OpenedAt   sql.NullString
}

type NewNotificationFromKafka struct {
	NotificationRepository entity.NotificationRepositoryInterface
}

func (n *NewNotificationFromKafka) Execute(input NotificationInputDTO) (*NotificationOutputDTO, error) {
	notification, err := entity.NewNotification(input.Title, input.Message)
	if err != nil {
		return nil, err
	}

	insertedId, err := n.NotificationRepository.Save(notification)
	if err != nil {
		return nil, err
	}

	return &NotificationOutputDTO{
		ID:         insertedId,
		Title:      notification.Title,
		Message:    notification.Message,
		ReceivedAt: notification.ReceivedAt,
		OpenedAt:   sql.NullString{},
	}, nil

}

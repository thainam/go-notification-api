package usecase

import (
	"go-notification-api/internal/entity"
)

type ListNotifications struct {
	NotificationRepository entity.NotificationRepositoryInterface
}

func (n *ListNotifications) Execute() ([]*entity.Notification, error) {
	notifications, err := n.NotificationRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return notifications, nil
}

package entity

type NotificationRepositoryInterface interface {
	Save(notification *Notification) (int, error)
	FindAll() ([]*Notification, error)
	FindById(notificationId int) (*Notification, error)
	SetOpenedAt(notification *Notification) (*Notification, error)
}

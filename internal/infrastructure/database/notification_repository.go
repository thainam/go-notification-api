package database

import (
	"database/sql"
	"errors"
	"go-notification-api/internal/entity"
	"go-notification-api/internal/logs"
	"time"
)

type NotificationRepository struct {
	Db *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{Db: db}
}

func (r *NotificationRepository) Save(notification *entity.Notification) (int, error) {
	result, err := r.Db.Exec("INSERT INTO notifications (title, message, received_at) VALUES (?, ?, ?)",
		notification.Title, notification.Message, notification.ReceivedAt)
	if err != nil {
		logs.DbLogError(err)
		return 0, errors.New("error while creating notification")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logs.DbLogError(err)
		return 0, errors.New("error while retrieving notification id")
	}
	return int(id), nil
}

func (r *NotificationRepository) FindAll() ([]*entity.Notification, error) {

	rows, err := r.Db.Query("SELECT id, title, message, received_at, opened_at FROM notifications ORDER BY received_at DESC")
	if err != nil {
		logs.DbLogError(err)
		return nil, errors.New("error while retrieving notifications")
	}

	var notifications []*entity.Notification
	for rows.Next() {
		var notification entity.Notification
		err = rows.Scan(&notification.ID, &notification.Title, &notification.Message, &notification.ReceivedAt, &notification.OpenedAt)
		if err != nil {
			logs.DbLogError(err)
			return nil, errors.New("error while retrieving notifications")
		}
		notifications = append(notifications, &notification)
	}
	return notifications, nil
}

func (r *NotificationRepository) FindById(notificationId int) (*entity.Notification, error) {
	var notification = entity.Notification{}
	err := r.Db.QueryRow("SELECT id, title, message, received_at, opened_at FROM notifications WHERE id = ?",
		notificationId).Scan(&notification.ID, &notification.Title, &notification.Message, &notification.ReceivedAt, &notification.OpenedAt)

	if err != nil {
		logs.DbLogError(err)
		return nil, errors.New("error while retrieving notification data")
	}

	return &notification, nil
}

func (r *NotificationRepository) SetOpenedAt(notification *entity.Notification) (*entity.Notification, error) {
	_, err := time.Parse("2006-01-02", notification.OpenedAt.String)
	if err == nil {
		return nil, nil
	}
	currentDateTime := time.Now().Format("2006-01-02")
	_, err = r.Db.Exec("UPDATE notifications SET opened_at = ? WHERE id = ?",
		currentDateTime, notification.ID)

	if err != nil {
		logs.DbLogError(err)
		return nil, errors.New("error while updating the date of notification opening")
	}

	notification.OpenedAt.String = currentDateTime

	return notification, nil
}

package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
	"go-notification-api/internal/entity"
	"go-notification-api/internal/logs"
	"testing"
)

type NotificationRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

func (suite *NotificationRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	_, err = db.Exec("CREATE TABLE `notifications` ( id INTEGER PRIMARY KEY AUTOINCREMENT, title VARCHAR(100) NOT NULL, message VARCHAR(200) NOT NULL, received_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, opened_at TIMESTAMP NULL DEFAULT NULL)")
	logs.DbLogError(err)
	suite.NoError(err)
	suite.Db = db
}

func (suite *NotificationRepositoryTestSuite) TearDownSuite() {
	err := suite.Db.Close()
	if err != nil {
		logs.DbLogError(err)
		return
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(NotificationRepositoryTestSuite))
}

func (suite *NotificationRepositoryTestSuite) TestNotificationRepository_Save() {
	notification, err := entity.NewNotification("Testing Title", "Testing Message")
	suite.NoError(err)

	repo := NewNotificationRepository(suite.Db)
	_, err = repo.Save(notification)
	suite.NoError(err)
}

func (suite *NotificationRepositoryTestSuite) TestNotificationRepository_FindAll() {

	notification, err := entity.NewNotification("Testing Title", "Testing Message")
	suite.NoError(err)

	repo := NewNotificationRepository(suite.Db)
	lastInsertId, err := repo.Save(notification)
	suite.NoError(err)

	queryResult, err := repo.FindAll()
	suite.NoError(err)

	suite.Equal(lastInsertId, queryResult[0].ID)
	suite.NotEmpty(queryResult[0].Title)
	suite.NotEmpty(queryResult[0].Message)
	suite.NotEmpty(queryResult[0].ReceivedAt)
}

func (suite *NotificationRepositoryTestSuite) TestNotificationRepository_FindById() {

	notification, err := entity.NewNotification("Testing Title", "Testing Message")
	suite.NoError(err)

	repo := NewNotificationRepository(suite.Db)
	id, err := repo.Save(notification)
	suite.NoError(err)

	queryResult, err := repo.FindById(id)
	suite.NoError(err)

	suite.NotEmpty(queryResult.ID)
	suite.NotEmpty(queryResult.Title)
	suite.NotEmpty(queryResult.Message)
	suite.NotEmpty(queryResult.ReceivedAt)
}

func (suite *NotificationRepositoryTestSuite) TestNotificationRepository_SetOpenedAt() {

	notification, err := entity.NewNotification("Testing Title", "Testing Message")
	suite.NoError(err)

	repo := NewNotificationRepository(suite.Db)
	id, err := repo.Save(notification)
	suite.NoError(err)

	notification, err = repo.FindById(id)
	suite.NoError(err)

	openedNotification, err := repo.SetOpenedAt(notification)

	suite.NotEmpty(openedNotification.OpenedAt.String)
	suite.NotEmpty(openedNotification.ID)
}

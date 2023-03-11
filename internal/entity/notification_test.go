package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotificationValidateShouldFailWhenTitleIsEmpty(t *testing.T) {
	var notification = Notification{}
	assert.Error(t, notification.Validate(), "title can't be empty")
}

func TestNotificationValidateShouldFailWhenMessageIsEmpty(t *testing.T) {
	var notification = Notification{Title: "Test"}
	assert.Error(t, notification.Validate(), "message can't be empty")
}

func TestNotificationValidateShouldSucceedWhenInformationIsValid(t *testing.T) {
	var notification = Notification{Title: "Test", Message: "Some Message"}
	assert.Nil(t, notification.Validate())
}

package notifications

import (
	"errors"
	"time"
)

type NotificationMethod interface {
	Send(text, sender string, time time.Time) int
}
type NotificationModule struct {
	notificationsInfo  map[int]InfoNotification
	notificationMethod NotificationMethod
}

func NewNotificationModule(method NotificationMethod) *NotificationModule {
	return &NotificationModule{
		notificationMethod: method,
		notificationsInfo:  make(map[int]InfoNotification),
	}
}
func (n *NotificationModule) Send(text, sender string) int {
	now := time.Now()
	id := n.notificationMethod.Send(text, sender, now)
	sendInfo := InfoNotification{
		Text:     text,
		Sender:   sender,
		TimeSend: now,
	}
	n.notificationsInfo[id] = sendInfo
	return id
}
func (n *NotificationModule) Info(id int) (InfoNotification, error) {
	info, ok := n.notificationsInfo[id]
	if !ok {
		return InfoNotification{}, errors.New("non-existent object")
	}
	return info, nil
}
func (n *NotificationModule) AllInfo() map[int]InfoNotification {
	tempMap := make(map[int]InfoNotification, len(n.notificationsInfo))
	for k, v := range n.notificationsInfo {
		tempMap[k] = v
	}
	return tempMap
}

package notifications

import "time"

type InfoNotification struct {
	Sender   string
	Text     string
	TimeSend time.Time
}

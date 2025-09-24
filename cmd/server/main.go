package main

import (
	"fmt"
	"github.com/k0kubun/pp"
	"mine/notifications"
	"mine/notifications/methods"
)

func main() {
	method := methods.NewPush()

	notificationModule := notifications.NewNotificationModule(method)
	idVlad := notificationModule.Send("Привет!", "Влад")
	notificationModule.Send("Hi!", "Alex")
	notificationModule.Info(idVlad)
	infoVlad, err := notificationModule.Info(idVlad)
	if err != nil {
		fmt.Println(err)
		return
	}
	pp.Println(infoVlad)
	allInfo := notificationModule.AllInfo()
	pp.Println(allInfo)
}

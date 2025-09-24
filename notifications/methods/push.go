package methods

import (
	"fmt"
	"math/rand"
	"time"
)

type Push struct {
}

func NewPush() Push {
	return Push{}
}

func (p Push) Send(text, sender string, time time.Time) int {
	fmt.Println("Пришло PUSH-уведомление")
	fmt.Printf("Время отправки: %v\nОтправитель: %s\nТекст уведомления: %s\n", time.Format("15:04 02.01.2006"), sender, text)
	fmt.Println("-----------------------------------")
	return rand.Int()
}

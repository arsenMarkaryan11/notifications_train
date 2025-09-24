package methods

import (
	"fmt"
	"math/rand"
	"time"
)

type SMS struct {
}

func NewSMS() SMS {
	return SMS{}
}

func (s SMS) Send(text, sender string, time time.Time) int {
	fmt.Println("Пришло SMS")
	fmt.Printf("Время отправки: %v\nОтправитель: %s\nТекст сообщения: %s\n", time.Format("15:04 02.01.2006"), sender, text)
	fmt.Println("-----------------------------------")

	return rand.Int()
}

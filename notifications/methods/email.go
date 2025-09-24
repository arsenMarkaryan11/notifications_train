package methods

import (
	"fmt"
	"math/rand"
	"time"
)

type Email struct {
}

func NewEmail() Email {
	return Email{}
}

func (e Email) Send(text, sender string, time time.Time) int {
	fmt.Println("Пришло сообщение через E-mail")
	fmt.Printf("Время отправки: %v\nОтправитель: %s\nТекст сообщения: %s\n", time.Format("15:04 02.01.2006"), sender, text)
	fmt.Println("-----------------------------------")
	return rand.Int()
}

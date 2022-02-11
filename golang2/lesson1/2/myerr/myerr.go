package myerr

import (
	"fmt"
	"time"
)

type myError struct {
	time time.Time
	text string
}

func (e *myError) Error() string {
	return fmt.Sprintf("[%v] Ошибка: %s\n", e.time, e.text)
}

func New(text string) *myError {
	return &myError{
		time: time.Now(),
		text: text,
	}
}

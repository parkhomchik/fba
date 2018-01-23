package model

import (
	"fmt"
)

//Error Ошибка
type Error struct {
	Error string `json:"error"`
}

//NewError создание "Ошибки"
func NewError(msg string) *Error {
	fmt.Println(msg)
	return &Error{Error: msg}
}

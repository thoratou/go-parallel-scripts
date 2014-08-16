package gsd

import (
	"fmt"
)

type Error struct {
	reason string
}

func NewError(args ...interface{}) *Error {
	value := fmt.Sprint("", args)
	return &Error{value}
}

func (self Error) Error() string {
	return self.reason
}

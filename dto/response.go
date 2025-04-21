package dto

import "time"

type Response[T any] struct {
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	Data      T         `json:"data"`
}

func NewResponseMessage(message string) Response[string] {
	return Response[string]{
		Timestamp: time.Now(),
		Message:   message,
	}
}

func NewResponseData[T any](data T) Response[T] {
	return Response[T]{
		Timestamp: time.Now(),
		Message:   "success",
		Data:      data,
	}
}

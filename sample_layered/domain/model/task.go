package model

import "time"

type Task struct {
	Id          int64
	Title       string
	Description string
	createdAt   time.Time
	updatedAt   time.Time
}

type createTaskRequest struct {
	Title       string
	Description string
	createdAt   time.Time
	updatedAt   time.Time
}

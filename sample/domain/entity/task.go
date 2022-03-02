package entity

import "time"

type Task struct {
	Id          int64
	Title       string
	Description string
	createdAt   time.Time
	updatedAt   time.Time
}

package model

type Task struct {
	ID          int64
	Title       string
	Description string
}

type TaskRequest struct {
	Title       string
	Description string
}

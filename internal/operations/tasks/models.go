package tasks

import (
	"time"
)

type Task struct {
	TaskId         int       `json:"task_id"`
	RegistrationId int       `json:"registration_id"`
	TaskDetail     string    `json:"task_detail"`
	TaskStatus     string    `json:"task_status"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
}

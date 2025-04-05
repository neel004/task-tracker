package storage

import (
	"time"
	taskModels "github.com/neel004/task-tracker/models"
)

type TaskItem struct{
	Id uint16	`json:id`
	Description string	`json:"description"`
	Status taskModels.TaskStatus	`json:"status"`
	CreatedAt time.Time	`json:"createdAt"`
	UpdatedAt time.Time	`json:"updatedAt"`
}

type Storage interface{
	Read() ([]TaskItem, error)
	Update([]TaskItem) error
}

package commands

import (
	"fmt"
	"time"
	"strings"
	fileStorage "github.com/neel004/task-tracker/storage"
	taskModels "github.com/neel004/task-tracker/models"
)

func Add(storage fileStorage.Storage, args ...string) (err error){
	items, err := storage.Read()
	if err != nil {
		fmt.Println("error encountered while reading storage: %w", err)
	}
	var new_id uint16
	if len(items) <= 0 {
		new_id = 1
	}else{
		new_id = items[len(items)-1].Id + 1
	}
	item := fileStorage.TaskItem{
		Id :  new_id,
		Description: strings.Join(args, " "),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status: taskModels.TODO,
	}

	items = append(items, item)
	if err = storage.Update(items); err != nil{
		return fmt.Errorf("error encountered while saving data: %w", err)
	}
	fmt.Printf("Task added successfully (ID: %d)\n", new_id)
	return nil
}

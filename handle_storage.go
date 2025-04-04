package main

import (
	"fmt"
	"os"
	"encoding/json"
	"time"
)

type TaskStatus int

const (
	TODO TaskStatus = iota // 0
	InProgress 
	Done
)
var MapFromString = map[string]TaskStatus{
	"todo": TODO,
	"inprogress": InProgress,
	"done": Done,
}
func ParseStatusType(input string) (TaskStatus, bool){
	taskStatus, ok := MapFromString[input]
	return taskStatus, ok
}


type TaskItem struct{
	Id uint16	`json:id`
	Description string	`json:"description"`
	Status TaskStatus	`json:"status"`
	CreatedAt time.Time	`json:"createdAt"`
	UpdatedAt time.Time	`json:"updatedAt"`
}

const dataDir string = "data"
const dataFile string = "data/tasks.json"

func ensureStorage() error {
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.Mkdir(dataDir, 0760); err != nil {
			return fmt.Errorf("error creating data directory: %w", err)
		}
	}

	if _, err := os.Stat(dataFile); err == nil {
		return nil
	}

	emptyJson := []TaskItem{}
	jsonData, err := json.Marshal(emptyJson)
	if err != nil {
		return fmt.Errorf("error encountered while creating empty json. %w", err)
	}
	if err := os.WriteFile(dataFile, jsonData, 0660); err != nil {
		return fmt.Errorf("error writing JSON file: %w", err)
	}
	return nil
}

func ReadStorage() ([]TaskItem, error){
	if err := ensureStorage(); err != nil {
		return nil, fmt.Errorf("storage setup error: %w", err)
	}

	itemByte, err := os.ReadFile(dataFile)
	if err != nil{
		return nil, fmt.Errorf("error encountered while reading file, %w", err)
	}

	var items []TaskItem
	err = json.Unmarshal(itemByte, &items)
	if err != nil{
		return nil, fmt.Errorf("error encountered while un-marshalling file contents. %w", err)
	}
	return items, nil
}

func UpdateStorage(items []TaskItem) error {
	err := ensureStorage()
	if err != nil{
		return fmt.Errorf("error encountered while ensuring storage: %w", err)
	}

	jsonData, err := json.Marshal(items)
	if err != nil{
		return fmt.Errorf("error encountered while marshalling the json. %w", err)
	}

	fileData, err := os.OpenFile(dataFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0660)
	if err != nil{
		return fmt.Errorf("error encountered while opening json. %w", err)
	}
	defer fileData.Close()

	_, err = fileData.Write(jsonData)
	if err != nil{
		return fmt.Errorf("error encountered while writing json. %w", err)
	}
	return nil
}

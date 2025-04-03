package main

import (
	"fmt"
	"time"
	"strings"
)

func Add(args ...string) (err error){
	items, err := ReadStorage()
	if err != nil {
		fmt.Println("error encountered while reading storage. %w", err)
	}
	var last_id uint16
	if len(items) <= 0 {
		last_id = 0
	}else{
		last_id = items[len(items)-1].Id
	}
	item := TaskItem{
		Id :  last_id + 1,
		Description: strings.Join(args, " "),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status: TODO,
	}

	items = append(items, item)
	if err = UpdateStorage(items); err != nil{
		return fmt.Errorf("error encountered while saving data. %w", err)
	}
	return nil
}

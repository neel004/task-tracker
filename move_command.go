

package main

import (
	"fmt"
	"time"
	"strconv"
)
func MoveTo(args ...string) error {
	items, err := ReadStorage()

	if err != nil {
		fmt.Println("error encountered while reading storage. %w", err)
	}
	if len(args) < 2{
		return fmt.Errorf("id and status needs to be passed for update.")
	}
	id, err := strconv.ParseUint(args[1], 10, 16)
	uint_16_id := uint16(id)
	if err != nil {
		return fmt.Errorf("error encountered while converting input to valid type, %w", err)
	}
	newState, ok := ParseStatusType(args[0]);
	if !ok {
		return fmt.Errorf("the queried status does not exists.")
	}
	var exists bool
	for idx, item := range items{
		if item.Id == uint_16_id{
			exists = true
			items[idx].Status = newState
			items[idx].UpdatedAt = time.Now()
		}
	}

	if !exists{
		return fmt.Errorf("task with id %d does not exists. please try again.", id)
	}

	if err = UpdateStorage(items); err != nil{
		return fmt.Errorf("error encountered while saving data. %w", err)
	}
	return nil
}

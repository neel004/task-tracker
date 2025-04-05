

package main

import (
	"fmt"
	"time"
	"strconv"
	"strings"
)
func MoveTo(args ...string) error {
	items, err := ReadStorage()
	if err != nil {
		return fmt.Errorf("error encountered while reading storage: %w", err)
	}

	if len(args) < 2{
		return fmt.Errorf("id and status needs to be passed for update.")
	}
	id, err := strconv.ParseUint(args[0], 10, 16)
	if err != nil {
		return fmt.Errorf("error encountered while converting input to valid type: %w", err)
	}
	uint_16_id := uint16(id)

	newState, ok := ParseStatusType(strings.ToLower(args[1]))
	if !ok {
		return fmt.Errorf("the queried status does not exist.")
	}

	var found bool
	for idx, item := range items{
		if item.Id == uint_16_id{
			found = true
			items[idx].Status = newState
			items[idx].UpdatedAt = time.Now()
			break
		}
	}

	if !found{
		return fmt.Errorf("task with id %d does not exists. please try again.", id)
	}

	if err = UpdateStorage(items); err != nil{
		return fmt.Errorf("error encountered while saving data: %w", err)
	}
	return nil
}

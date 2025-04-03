package main

import (
	"fmt"
	"strconv"
	"slices"
)

func Delete(args ... string) error {
	items, err := ReadStorage()

	if err != nil {
		fmt.Println("error encountered while reading storage. %w", err)
	}
	if len(args) < 1{
		return fmt.Errorf("id and description needs to be passed for update.")
	}
	id, err := strconv.ParseUint(args[0], 10, 16)
	uint_16_id := uint16(id)
	if err != nil {
		return fmt.Errorf("error encountered while converting input to valid type, %w", err)
	}

	index := slices.IndexFunc(items, func(item TaskItem) bool {
		return item.Id == uint_16_id
	})
	if index < 0 {
		return fmt.Errorf("item with queried id is not present.")
	}
	
	items = slices.Delete(items, index, index+1)

	if err = UpdateStorage(items); err != nil{
		return fmt.Errorf("error encountered while saving data. %w", err)
	}
	return nil
}

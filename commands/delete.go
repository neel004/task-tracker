package commands

import (
	"fmt"
	fileStorage "github.com/neel004/task-tracker/storage"
	"slices"
	"strconv"
)

func Delete(storage fileStorage.Storage, args ...string) error {
	items, err := storage.Read()

	if err != nil {
		fmt.Println("error encountered while reading storage: %w", err)
	}
	if len(args) < 1 {
		return fmt.Errorf("id and description needs to be passed for update.")
	}
	id, err := strconv.ParseUint(args[0], 10, 16)
	uint_16_id := uint16(id)
	if err != nil {
		return fmt.Errorf("error encountered while converting input to valid type: %w", err)
	}

	index := slices.IndexFunc(items, func(item fileStorage.TaskItem) bool {
		return item.Id == uint_16_id
	})
	if index < 0 {
		return fmt.Errorf("item with queried id is not present.")
	}

	items = slices.Delete(items, index, index+1)

	if err = storage.Update(items); err != nil {
		return fmt.Errorf("error encountered while saving data: %w", err)
	}
	fmt.Printf("Task deleted successfully (ID: %d)\n", uint_16_id)
	return nil
}

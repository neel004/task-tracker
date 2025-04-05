package commands

import (
	"fmt"
	"time"
	"strings"
	fileStorage "github.com/neel004/task-tracker/storage"
)

func wrapText(width int, input string) []string{
	words := strings.Fields(input)
	lines := []string{}
	line := ""

	for _, word := range words{
		if len(line) + len(word) > width{
			lines = append(lines, line)
			line = word
		}else{
			line += word + " "
		}
	}
	if len(line) > 0{
		lines = append(lines, line)
	}
	if len(lines) == 0{
		lines = append(lines, "")
	}
	return lines
}

func List(storage fileStorage.Storage, args ...string) error{
	items, err := storage.Read()

	if err != nil {
		fmt.Println("error encountered while reading storage: %w", err)
	}
	var status string
	if len(args) > 0{
		status = args[0]
	}
	fmt.Printf("%-3s %-20s %-10s %-20s\n", "Id", "Description", "Status", "Last Updated")
	fmt.Println(strings.Repeat("-", 55))

	for _, item := range items{
		if status == "" || strings.ToUpper(item.Status.String()) == strings.ToUpper(status){
			lines := wrapText(20, item.Description)
			fmt.Printf("%-3d %-20s %-10s %-20s\n", item.Id, lines[0], item.Status, item.UpdatedAt.Format(time.RFC822))
			for _, line := range lines[1:]{
				fmt.Printf("%-3s %-20s \n", "", line)
			}
		}
	}
	return nil
}

package main

import (
	"fmt"
	"github.com/neel004/task-tracker/commands"
	"github.com/neel004/task-tracker/storage"
	"os"
)

type Commands struct {
	name     string
	Desc     string
	callback func(storage.Storage, ...string) (err error)
}

var CommandsMap = map[string]Commands{
	"add":    {name: "add", Desc: "Add to-do items in list. ex. add <description>", callback: commands.Add},
	"update": {name: "update", Desc: "Update any item's description using id. ex. update 1 <newdescription>", callback: commands.Update},
	"delete": {name: "delete", Desc: "Delete any item using id. ex. delete 1", callback: commands.Delete},
	"move":   {name: "move", Desc: "Move any tasks to inprogress, todo, or done state using id. ex. move todo 1", callback: commands.MoveTo},
	"list":   {name: "list", Desc: "List tasks. ex. list (lists all), list todo(lists task with todo state)", callback: commands.List},
}

// Function to display help
func showHelp() {
	fmt.Println("Usage instructions are as follows:")
	for key, cmd := range CommandsMap {
		fmt.Printf("%-10s %s\n", key, cmd.Desc)
	}
}

func main() {
	inputs := os.Args[1:]
	if len(inputs) < 1 || inputs[0] == "help" {
		showHelp()
		return
	}

	command_name := inputs[0]
	val, ok := CommandsMap[command_name]
	if !ok {
		err := fmt.Errorf("the requested command is not supported.")
		fmt.Println(err.Error())
		return
	}
	storage := storage.GetStorage()
	err := val.callback(storage, inputs[1:]...)
	if err != nil {
		fmt.Println(err.Error())
	}

}

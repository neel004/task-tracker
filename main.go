package main


import (
	"fmt"
	"bufio"
	"os"
	"strings"
)
type commands struct{
	name string
	Desc string
	callback func(...string)(err error)
}
func getCommands() (map[string]commands) {
	return map[string]commands{
		"add" : {name: "add", Desc:"Add to-do items in list. ex. add <description>", callback:Add},
		"update": {name: "update", Desc: "Update any item's description using id. ex. update 1 <newdescription>", callback: Update},
		"delete": {name: "delete", Desc: "Delete any item using id. ex. delete 1", callback: Delete},
		"move": {name: "inProgress", Desc: "Move any tasks to in-progress, todo, or done state using id. ex. move todo 1", callback: MoveTo},
		"list": {name: "list", Desc: "List tasks. ex. list (lists all), list todo(lists task with todo state)", callback: List},
	}
}

func main(){
	commands := getCommands()
	reader := bufio.NewReader(os.Stdin)
	line, _, err := reader.ReadLine()
	
	if err != nil{
		err = fmt.Errorf("error encountered while reading line from stdin.")
		fmt.Println(err)
		return
	}

	inputs := strings.Split(strings.TrimSpace(string(line)), " ")
	fmt.Println(inputs, len(inputs))
	if len(inputs) < 2{
		err = fmt.Errorf("The input is invalid please check again. or use help.")
		fmt.Println(err.Error())
		return
	}
	command_name := inputs[1]
	val, ok := commands[command_name]
	fmt.Println(command_name == "help")
	if command_name == "help" {
		fmt.Println("Usage instructions are as follows:")
		for key, val := range getCommands() {
			fmt.Printf("%-6s %-20s\n", key, val.Desc)
		}
		return
	}else if !ok{
		err:= fmt.Errorf("the requested command is not supported.")
		fmt.Println(err.Error())
		return
	}else{
		err := val.callback(inputs[2:]...)
		if err != nil{
			fmt.Println(err.Error())
		}
	}
}

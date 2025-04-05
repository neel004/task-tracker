package models

type TaskStatus int

const (
	TODO TaskStatus = iota // 0
	InProgress
	Done
)

var MapFromString = map[string]TaskStatus{
	"todo":       TODO,
	"inprogress": InProgress,
	"done":       Done,
}

func ParseStatusType(input string) (TaskStatus, bool) {
	taskStatus, ok := MapFromString[input]
	return taskStatus, ok
}

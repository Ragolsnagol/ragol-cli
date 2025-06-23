package main

type ToDo struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func NewToDo(task string) *ToDo {
	return &ToDo{Task: task, Done: false}
}

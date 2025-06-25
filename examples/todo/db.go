package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// This isn't really a database of any kind it's just a bunch of functions for saving to json
// and doing actions on the json

const fileName string = "todos.json"

func AddToJson(todo ToDo) error {
	todos, err := GetFromJson()
	if err != nil {
		return err
	}

	todos = append(todos, todo)

	err = saveTodoJson(todos)
	if err != nil {
		return err
	}

	return nil
}

func GetFromJson() ([]ToDo, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []ToDo{}, nil
		}
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var data []ToDo
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return data, nil
}

func DeleteFromJson(t string) error {
	todos, err := GetFromJson()
	if err != nil {
		return err
	}

	for i, todo := range todos {
		if strings.EqualFold(todo.Task, t) {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}

	err = saveTodoJson(todos)
	if err != nil {
		return err
	}

	return nil
}

func MarkDoneFromJson(t string) error {
	todos, err := GetFromJson()
	if err != nil {
		return err
	}

	for i, todo := range todos {
		if strings.EqualFold(todo.Task, t) {
			todo.Done = true
			todos[i] = todo
			break
		}
	}

	err = saveTodoJson(todos)
	if err != nil {
		return err
	}

	return nil
}

func saveTodoJson(todos []ToDo) error {
	jsonData, err := json.MarshalIndent(todos, "", "    ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}

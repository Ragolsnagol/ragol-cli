package main

import (
	"fmt"
	"ragol-cli/core"
	"ragol-cli/core/action"
	"ragol-cli/core/command"
	"ragol-cli/core/context"
	"ragol-cli/core/flag"
)

func main() {
	app := core.NewApp(
		"ToDo cli",
		"1.0",
		[]command.BaseCommand{
			addTodoCommand(),
			deleteTodoCommand(),
			markDoneCommand(),
			getToDosCommand(),
		})
	err := app.Run()
	if err != nil {
		panic(err)
	}
}

func addTodoCommand() command.BaseCommand {
	c := command.NewCommand(
		"add",
		"Add a new todo",
		addTodoAction(),
		[]flag.Flag{
			*flag.NewFlag("task", "t", true, true),
		},
	)
	return *c
}

func addTodoAction() action.Action {
	a := action.NewAction(func(ctx context.Context) error {
		task, err := ctx.GetFlag("task")
		if err != nil {
			return err
		}

		todo := NewToDo(task.Value.(string))

		err = AddToJson(*todo)
		if err != nil {
			return err
		}

		return nil
	})
	return a
}

func deleteTodoCommand() command.BaseCommand {
	c := command.NewCommand(
		"delete",
		"Delete a todo using the task name",
		deleteTodoAction(),
		[]flag.Flag{
			*flag.NewFlag("task", "t", true, true),
		},
	)

	return *c
}

func deleteTodoAction() action.Action {
	a := action.NewAction(func(ctx context.Context) error {
		task, err := ctx.GetFlag("task")
		if err != nil {
			return err
		}

		err = DeleteFromJson(task.Value.(string))
		if err != nil {
			return err
		}

		return nil
	})
	return a
}

func markDoneCommand() command.BaseCommand {
	c := command.NewCommand(
		"done",
		"Mark a task as done",
		markDoneAction(),
		[]flag.Flag{
			*flag.NewFlag("task", "t", true, true),
		},
	)

	return *c
}

func markDoneAction() action.Action {
	a := action.NewAction(func(ctx context.Context) error {
		task, err := ctx.GetFlag("task")
		if err != nil {
			return err
		}

		err = MarkDoneFromJson(task.Value.(string))
		if err != nil {
			return err
		}

		return nil
	})

	return a
}

func getToDosCommand() command.BaseCommand {
	c := command.NewCommand(
		"get",
		"Get all todos",
		getToDosAction(),
		[]flag.Flag{},
	)

	return *c
}

func getToDosAction() action.Action {
	a := action.NewAction(func(ctx context.Context) error {
		todos, err := GetFromJson()
		if err != nil {
			return err
		}

		fmt.Println("ToDos:")
		for i, todo := range todos {
			fmt.Printf("%v: %v", i+1, todo.Task)
		}

		return nil
	})

	return a
}

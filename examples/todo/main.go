package main

import (
	"ragol-cli/core"
	"ragol-cli/core/action"
	"ragol-cli/core/command"
	"ragol-cli/core/context"
	"ragol-cli/core/flag"
)

func main() {
	app := core.NewApp(
		"",
		"1.0",
		[]command.BaseCommand{
			addTodoCommand(),
			deleteTodoCommand(),
			markDoneCommand(),
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

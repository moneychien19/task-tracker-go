package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/moneychien19/task-tracker-go/task"
)

func main() {
	app := &cli.App{
		Name:  "task tracker",
		Usage: "manage tasks with CLI",
		Action: func(cCtx *cli.Context) error {
			command := cCtx.Args().Get(0)

			switch command {
			case "add":
				id, err := strconv.ParseInt(cCtx.Args().Get(1), 10, 64)
				if err != nil {
					return err
				}
				description := cCtx.Args().Get(2)
	
				err = task.AddTaskToJson(id, description)
				if err == nil {
					fmt.Printf("Task added successfully (ID: %d)\n", id)
				}
			case "update":
				id, err := strconv.ParseInt(cCtx.Args().Get(1), 10, 64)
				if err != nil {
					return err
				}
				description := cCtx.Args().Get(2)
				err = task.UpdateTaskToJson(id, description)
				if err == nil {
					fmt.Printf("Task updated successfully (ID: %d)\n", id)
				}
			case "delete":
				id, err := strconv.ParseInt(cCtx.Args().Get(1), 10, 64)
				if err != nil {
					return err
				}
				err = task.DeleteTaskFromJson(id)
				if err == nil {
					fmt.Printf("Task deleted successfully (ID: %d)\n", id)
				}
			case "mark-in-progress":
				id, err := strconv.ParseInt(cCtx.Args().Get(1), 10, 64)
				if err != nil {
					return err
				}
				err = task.ChangeTaskStatus(id, task.StatusInProgress)
				if err == nil {
					fmt.Printf("Task marked as in-progress successfully (ID: %d)\n", id)
				}
			case "mark-done":
				id, err := strconv.ParseInt(cCtx.Args().Get(1), 10, 64)
				if err != nil {
					return err
				}
				err = task.ChangeTaskStatus(id, task.StatusDone)
				if err == nil {
					fmt.Printf("Task marked as done successfully (ID: %d)\n", id)
				}
			case "list":
				status := cCtx.Args().Get(1)
				tasks, err := task.GetTasks(status)
				if err != nil {
					return err
				}
				task.RenderTaskTables(tasks)
			}				
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Something went wrong!")
	}
}

package task_test

import (
	"os"
	"testing"

	"github.com/moneychien19/task-tracker-go/task"
)

func setup() {
	os.Create("tasks.json")
}

func teardown() {
	os.Remove("tasks.json")
}

func TestAddTaskToJson(t *testing.T) {
	setup()
	defer teardown()

	err := task.AddTaskToJson(1, "Test Task")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, err := task.GetTasks("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Description != "Test Task" {
		t.Fatalf("Expected 'Test Task', got %s", tasks[0].Description)
	}
}

func TestUpdateTaskToJson(t *testing.T) {
	setup()
	defer teardown()

	task.AddTaskToJson(1, "Test Task")
	err := task.UpdateTaskToJson(1, "Updated Task")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, err := task.GetTasks("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if tasks[0].Description != "Updated Task" {
		t.Fatalf("Expected 'Updated Task', got %s", tasks[0].Description)
	}
}

func TestDeleteTaskFromJson(t *testing.T) {
	setup()
	defer teardown()

	task.AddTaskToJson(1, "Test Task")
	err := task.DeleteTaskFromJson(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, err := task.GetTasks("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(tasks) != 0 {
		t.Fatalf("Expected 0 tasks, got %d", len(tasks))
	}
}

func TestChangeTaskStatus(t *testing.T) {
	setup()
	defer teardown()

	task.AddTaskToJson(1, "Test Task")
	err := task.ChangeTaskStatus(1, task.StatusDone)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	tasks, err := task.GetTasks("")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if tasks[0].Status != task.StatusDone {
		t.Fatalf("Expected status 'done', got %s", tasks[0].Status)
	}
}

func TestListTasks(t *testing.T) {
	setup()
	defer teardown()

	task.AddTaskToJson(1, "Test Task 1")
	task.AddTaskToJson(2, "Test Task 2")
	task.ChangeTaskStatus(2, task.StatusDone)

	tasks, err := task.GetTasks(task.StatusDone)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(tasks) != 1 {
		t.Fatalf("Expected 1 task, got %d", len(tasks))
	}

	if tasks[0].Description != "Test Task 2" {
		t.Fatalf("Expected 'Test Task 2', got %s", tasks[0].Description)
	}
}
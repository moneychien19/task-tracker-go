package task

import (
	"encoding/json"
	"os"
	"time"
)

const (
	StatusTodo = "todo"
	StatusInProgress = "in-progress"
	StatusDone = "done"
)

type Task struct {
	Id int64 `json:"id"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
func new(id int64, description string, status string, createdAt time.Time, updatedAt time.Time) Task {
	return Task{id, description, status, createdAt, updatedAt}
}

func readTasksFromJson() ([]Task, error) {
	_, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		_, err = os.Create("tasks.json")
		if err != nil {
			return nil, err
		}
	}

	jsonTasks, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, err
	}

	tasks := make([]Task, 0)
	if len(jsonTasks) > 0 {
		err = json.Unmarshal(jsonTasks, &tasks)
		if err != nil {
			return tasks, err
		}
	}
	return tasks, nil
}
func writeTasksToJson(tasks []Task) error {
	jsonTasks, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile("tasks.json", jsonTasks, 0644)
}

func AddTaskToJson(id int64, description string) (error) {
	existedTasks, err := readTasksFromJson()
	if err != nil {
		return err
	}
	
	task := new(id, description, StatusTodo, time.Now(), time.Now())
	updatedTasks := append(existedTasks, task)

	return writeTasksToJson(updatedTasks)
}

func UpdateTaskToJson(id int64, description string) (error) {
	existedTasks, err := readTasksFromJson()
	if err != nil {
		return err
	}

	for i, t := range existedTasks {
		if t.Id == id {
			existedTasks[i].Description = description
			existedTasks[i].UpdatedAt = time.Now()
			break
		}
	}

	return writeTasksToJson(existedTasks)
}

func DeleteTaskFromJson(id int64) (error) {
	existedTasks, err := readTasksFromJson()
	if err != nil {
		return err
	}

	for i, t := range existedTasks {
		if t.Id == id {
			existedTasks = append(existedTasks[:i], existedTasks[i+1:]...)
			break
		}
	}

	return writeTasksToJson(existedTasks)
}

func ChangeTaskStatus(id int64, status string) (error) {
	existedTasks, err := readTasksFromJson()
	if err != nil {
		return err
	}

	for i, t := range existedTasks {
		if t.Id == id {
			existedTasks[i].Status = status
			existedTasks[i].UpdatedAt = time.Now()
			break
		}
	}

	return writeTasksToJson(existedTasks)
}

func ListTasks(status string) ([]Task, error) {
	tasks, err := readTasksFromJson()
	if err != nil {
		return nil, err
	}

	if status != "" {
		filteredTasks := make([]Task, 0)
		for _, t := range tasks {
			if t.Status == status {
				filteredTasks = append(filteredTasks, t)
			}
		}
		return filteredTasks, nil
	}

	return tasks, err
}
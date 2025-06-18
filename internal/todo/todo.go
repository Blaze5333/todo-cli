package todo

import (
	"errors"
	"time"

	"github.com/Blaze5333/todo-cli/internal/storage"
)

type Todo struct {
	Username string `json:"username"`
	Task     []Task `json:"task"`
}
type Task struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Priority    int       `json:"priority"`
}

var strg = storage.NewStorage[[]Todo]("todos.json")

func validateTask(task []Task, index int) bool {
	if index < 0 || index >= len(task) {
		return false
	}
	return true
}

func AddTask(username, title, description string, priority int) (*Task, error) {
	todos := []Todo{}
	err := strg.Load(&todos)
	if err != nil {
		return nil, err
	}
	//what if the user does not exist?
	//what if the array is empty?
	newTask := Task{
		Title:       title,
		Description: description,
		Priority:    priority,
		Done:        false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	for i, todo := range todos {
		if todo.Username == username {
			todos[i].Task = append(todos[i].Task, newTask)
			err = strg.Save(todos)
			if err != nil {
				return nil, err
			}
			return &newTask, nil
		}
	}
	// If the user does not exist, create a new Todo entry
	newTodo := Todo{
		Username: username,
		Task:     []Task{newTask},
	}
	todos = append(todos, newTodo)
	err = strg.Save(todos)
	if err != nil {
		return nil, err
	}
	return &newTask, nil // User not found, return nil
}
func GetTasks(username string) ([]Task, error) {
	todos := []Todo{}
	err := strg.Load(&todos)
	if err != nil {
		return nil, err
	}
	for _, todo := range todos {
		if todo.Username == username {
			return todo.Task, nil
		}
	}
	return []Task{}, nil // User not found, return empty slice
}
func UpdateTask(username, description string, done bool, index int) error {
	todos := []Todo{}
	err := strg.Load(&todos)
	if err != nil {
		return err
	}
	for _, todo := range todos {
		if todo.Username == username {
			if !validateTask(todo.Task, index) {
				return errors.New("invalid task index")
			}
			todo.Task[index].Description = description
			todo.Task[index].UpdatedAt = time.Now()
			return strg.Save(todos)
		}
	}
	return errors.New("task not found") // Task not found, no error
}
func DeleteTask(username string, index int) error {
	todos := []Todo{}
	err := strg.Load(&todos)
	if err != nil {
		return err
	}
	for i, todo := range todos {
		if todo.Username == username {
			if !validateTask(todo.Task, index) {
				return errors.New("invalid task index")
			}
			todos[i].Task = append(todos[i].Task[:index], todos[i].Task[index+1:]...)
			return strg.Save(todos)
		}
	}
	return errors.New("task not found") // Task not found, no error
}
func CompleteTask(username string, index int) error {
	todos := []Todo{}
	err := strg.Load(&todos)
	if err != nil {
		return err
	}
	for _, todo := range todos {
		if todo.Username == username {
			if !validateTask(todo.Task, index) {
				return errors.New("invalid task index")
			}
			todo.Task[index].Done = true
			todo.Task[index].UpdatedAt = time.Now()
			return strg.Save(todos)
		}
	}
	return errors.New("task not found")
}

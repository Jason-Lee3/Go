package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	command := arguments[1]

	switch command {
	case "add":
		err := addTask(arguments[2])
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	case "list":
		listTask()
	case "delete":
		deleteTask(arguments[2])
	case "complete":
		completeTask(arguments[2])
	default:
		fmt.Println("Unknown command, exiting the program...")
		os.Exit(1)
	}
}

func loadTask() ([]Task, error) {
	initTask := []Task{}
	read, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return initTask, nil
		}
		return initTask, err
	}

	err = json.Unmarshal(read, &initTask)
	if err != nil {
		return initTask, err
	}
	return initTask, nil
}

func saveTask(tasks []Task) error {
	encode, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", encode, 0666)
	if err != nil {
		fmt.Println("Error occured while writing to file...")
		return err
	}

	return nil

}

func addTask(text string) error {
	tasks, err := loadTask()
	if err != nil {
		return err
	}
	max := 0
	for _, task := range tasks {
		if task.id > max {
			max = task.id
		}
	}
	newMaxId := max + 1

	task := Task{
		id:        newMaxId,
		text:      text,
		completed: false,
	}

	tasks = append(tasks, task)

	err = saveTask(tasks)
	if err != nil {
		return err
	}
	return nil

}

func listTask() {
	tasks, err := loadTask()
	if err != nil {
		os.Exit(1)
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks yet. Add one with: go run add \"Buy Milk\"")
	}

	for _, task := range tasks {
		fmt.Print(strconv.Itoa(task.id) + ".")
		if task.completed {
			fmt.Print(" [x] ")
		} else {
			fmt.Print(" [ ] ")
		}
		fmt.Println(task.text)
	}
}

func deleteTask(id string) {
	tasks, err := loadTask()
	if err != nil {
		os.Exit(1)
	}
	tempTasks := []Task{}
	ident, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("String cannot be converted to integer. Exiting...")
		os.Exit(1)
	}

	for _, task := range tasks {
		if task.id != ident {
			tempTasks = append(tempTasks, task)
		}
	}

	if len(tempTasks) == len(tasks) {
		fmt.Printf("Task %+v not found\n", id)
	} else {
		fmt.Printf("Deleted task %+v\n", id)
	}

	err = saveTask(tempTasks)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func completeTask(id string) {
	tasks, err := loadTask()
	newTasks := []Task{}
	if err != nil {
		os.Exit(1)
	}

	// convert string to int
	intId, err := strconv.Atoi(id)
	if err != nil {
		os.Exit(1)
	}
	for _, task := range tasks {
		if task.id == intId {
			task.completed = true
		}
		newTasks = append(newTasks, task)
	}

	err = saveTask(newTasks)
	if err != nil {
		os.Exit(1)
	}

}

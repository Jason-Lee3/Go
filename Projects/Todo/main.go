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

	decodeJsonError := json.Unmarshal(read, &initTask)
	if decodeJsonError != nil {
		return initTask, decodeJsonError
	}
	return initTask, nil
}

func saveTask(tasks []Task) error {
	encode, encodeJsonError := json.Marshal(tasks)
	if encodeJsonError != nil {
		return encodeJsonError
	}
	writeFileError := os.WriteFile("tasks.json", encode, 0666)
	if writeFileError != nil {
		fmt.Println("Error occured while writing to file...")
		return writeFileError
	}

	return nil

}

func addTask(text string) error {
	tasks, err := loadTask()
	if err != nil {
		return err
	}
	var max int = 0
	for _, task := range tasks {
		if task.Id > max {
			max = task.Id
		}
	}
	var newMaxId int = max + 1

	task := Task{
		Id:        newMaxId,
		Text:      text,
		Completed: false,
	}

	tasks = append(tasks, task)

	saveTaskError := saveTask(tasks)
	if saveTaskError != nil {
		return saveTaskError
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
		fmt.Print(strconv.Itoa(task.Id) + ".")
		if task.Completed {
			fmt.Print(" [x] ")
		} else {
			fmt.Print(" [ ] ")
		}
		fmt.Println(task.Text)
	}
}

func deleteTask(id string) {
	tasks, err := loadTask()
	if err != nil {
		os.Exit(1)
	}
	tempTasks := []Task{}
	ident, conversionError := strconv.Atoi(id)

	if conversionError != nil {
		fmt.Println("String cannot be converted to integer. Exiting...")
		os.Exit(1)
	}

	for _, task := range tasks {
		if task.Id != ident {
			tempTasks = append(tempTasks, task)
		}
	}

	if len(tempTasks) == len(tasks) {
		fmt.Printf("Task %+v not found\n", id)
	} else {
		fmt.Printf("Deleted task %+v\n", id)
	}

	saveError := saveTask(tempTasks)

	if saveError != nil {
		fmt.Println("Error:", saveError)
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
	intId, convertError := strconv.Atoi(id)
	if convertError != nil {
		os.Exit(1)
	}
	for _, task := range tasks {
		if task.Id == intId {
			task.Completed = true
		}
		newTasks = append(newTasks, task)
	}

	saveError := saveTask(newTasks)
	if saveError != nil {
		os.Exit(1)
	}

}

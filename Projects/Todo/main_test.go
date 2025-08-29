package main

import "testing"

func TestAddTask(t *testing.T) {
	tasks, loadTaskError := loadTask()
	if loadTaskError != nil {
		t.Errorf("Expected no errors")
	}

	err := addTask("Write Test Cases")
	if err != nil {
		t.Errorf("Expected successful addition to task")
	}

	resultTasks, loadTaskError := loadTask()
	if loadTaskError != nil {
		t.Error("Found error:", loadTaskError)
	}
	if len(tasks) == len(resultTasks) {
		t.Errorf("Different length expected. tasks length: %+v, resultTasks length: %+v", len(tasks), len(resultTasks))
	}
}

func TestLoadTask(t *testing.T) {

}

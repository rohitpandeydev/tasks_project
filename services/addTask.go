package services

import (
	"fmt"
	"time"

	"github.com/rohitpandeydev/tasks/types"
)

var filename = "./tasks.csv"

func WriteTask(description string) {
	tasks, err := ReadCSVFile(filename)
	if err != nil {
		fmt.Println("sorry there was some internal issue")
		return
	}
	task := types.Task{
		ID:          len(tasks) + 1,
		Description: description,
		CreatedAt:   time.Now(),
		IsComplete:  false,
	}
	tasks = append(tasks, &task)
	err = WriteCSVFile(filename, tasks)
	if err != nil {
		fmt.Println("sorry there was some internal issue")
	}
}

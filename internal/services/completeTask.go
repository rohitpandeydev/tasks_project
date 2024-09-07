package services

import (
	"fmt"
)

func CompleteTask(id int) {
	tasks, err := ReadCSVFile(filename)
	if err != nil {
		fmt.Println("some internal error")
		return
	}

	for _, task := range tasks {
		if task.ID == id {
			task.IsComplete = true
		}
	}
	err = WriteCSVFile(filename, tasks)
	if err != nil {
		fmt.Println("some internal error")
		return
	}
	fmt.Println("task completed successfully")
}

package services

import (
	"fmt"

	"github.com/rohitpandeydev/tasks/types"
)

func DeleteTask(id int) {

	tasks := []*types.Task{}
	oldTasks, err := ReadCSVFile(filename)
	if err != nil {
		fmt.Println("some internal error")
		return
	}

	for _, task := range oldTasks {
		if task.ID != id {
			tasks = append(tasks, task)
		}
	}
	err = WriteCSVFile(filename, tasks)
	if err != nil {
		fmt.Println("some internal error")
		return
	}
	fmt.Println("task deleted successfully")
}

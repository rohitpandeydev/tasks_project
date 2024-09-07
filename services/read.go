package services

import (
	"fmt"
	"os"

	"github.com/rohitpandeydev/tasks/types"
	"github.com/rohitpandeydev/tasks/utils"

	"github.com/gocarina/gocsv"
)

func ReadCSVFile(filename string) ([]*types.Task, error) {
	tasks := []*types.Task{}
	file, err := utils.LoadFile(filename)
	if err != nil {
		// If the file doesn't exist, return an empty slice
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, fmt.Errorf("error while reading the file: %w", err)
	}
	defer utils.CloseFile(file)

	if err := gocsv.UnmarshalFile(file, &tasks); err != nil {
		if err.Error() == "empty csv file given" {
			// If the file is empty, return an empty slice
			return tasks, nil
		}
		return nil, fmt.Errorf("file not properly formatted: %w", err)
	}

	return tasks, nil
}

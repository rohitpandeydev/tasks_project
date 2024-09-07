package services

import (
	"fmt"
	"os"

	"github.com/rohitpandeydev/tasks/internal/types"
	"github.com/rohitpandeydev/tasks/internal/utils"

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

func WriteCSVFile(filename string, tasks []*types.Task) error {
	file, err := utils.LoadFile(filename)
	if err != nil {
		return err
	}
	err = gocsv.MarshalFile(&tasks, file)
	if err != nil {
		return err
	}

	if err = utils.CloseFile(file); err != nil {
		return err
	}
	return nil
}

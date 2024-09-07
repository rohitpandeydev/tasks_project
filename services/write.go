package services

import (
	"log"

	"github.com/rohitpandeydev/tasks/types"
	"github.com/rohitpandeydev/tasks/utils"

	"github.com/gocarina/gocsv"
)

func WriteCSVFile(filename string, tasks []*types.Task) error {
	file, err := utils.LoadFile(filename)
	if err != nil {
		log.Fatal("Error while loading the file", err)
		return err
	}
	err = gocsv.MarshalFile(&tasks, file)
	if err != nil {
		log.Fatal("error whiel writing the file", err)
	}

	if err = utils.CloseFile(file); err != nil {
		log.Fatal("unable to close the file", err)
		return err
	}
	return nil
}

package services

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/mergestat/timediff"
)

func ListOutputDetailed(detail bool) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	tasks, err := ReadCSVFile(filename)
	if err != nil {
		if err.Error() == "empty csv file given" {
			fmt.Println("No tasks found")
		} else {
			fmt.Printf("Error reading the file: %v\n", err)
		}
		return
	}

	if detail {
		fmt.Fprintln(w, "ID\tDescription\tCreatedAt\tIsComplete")
	} else {
		fmt.Fprintln(w, "ID\tDescription\tCreatedAt")
	}

	// Print tasks

	for _, task := range tasks {
		if detail {
			fmt.Fprintf(w, "%d\t%s\t%s\t%v\n", task.ID, task.Description, timediff.TimeDiff(task.CreatedAt), task.IsComplete)
		} else if !task.IsComplete {
			fmt.Fprintf(w, "%d\t%s\t%s\n", task.ID, task.Description, timediff.TimeDiff(task.CreatedAt))
		} else {
			continue
		}
	}
	w.Flush()
}

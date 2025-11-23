package cmd

import (
	"fmt"

	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/ryanmalonzo/eat-that-frog/internal/frog"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark today's frog as done",
	Long:  `Mark the task that has been picked as today's frog to eat as done.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		frogTask, err := frog.GetTodayFrog()
		if err != nil {
			return err
		}

		err = db.MarkFrogAsDone(frogTask)
		if err != nil {
			return err
		}

		cmd.Println(fmt.Sprintf("✅ %s", frogTask))
		return nil
	},
}

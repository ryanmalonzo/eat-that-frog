package cmd

import (
	"fmt"

	"github.com/ryanmalonzo/eat-that-frog/internal/db"
	"github.com/ryanmalonzo/eat-that-frog/internal/frog"
	"github.com/spf13/cobra"
)

var skipCmd = &cobra.Command{
	Use:   "skip",
	Short: "Skip today's frog",
	Long:  `Skip the task that has been picked as today's frog to eat.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		frogTask, err := frog.GetTodayFrog()
		if err != nil {
			return err
		}

		err = db.SkipTodayFrog(frogTask)
		if err != nil {
			return err
		}

		cmd.Println(fmt.Sprintf("❌ %s", frogTask))
		return nil
	},
}

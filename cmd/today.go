package cmd

import (
	"github.com/ryanmalonzo/eat-that-frog/internal/frog"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show today's frog",
	Long:  `Show the task that has been picked as today's frog to eat.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		frog, err := frog.GetTodayFrog()
		if err != nil {
			return err
		}

		cmd.Println(frog)
		return nil
	},
}

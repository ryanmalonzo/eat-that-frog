package cmd

import (
	"github.com/ryanmalonzo/eat-that-frog-cli/internal/db"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Show today's frog",
	Long:  `Show the task that has been picked as today's frog to eat.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		frog, err := db.GetTodayFrog()
		if err != nil {
			return err
		}

		if frog == "" {
			cmd.Println("No frog has been picked for today yet.")
			return nil
		}

		cmd.Println(frog)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}

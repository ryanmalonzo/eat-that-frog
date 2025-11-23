package cmd

import (
	"github.com/ryanmalonzo/eat-that-frog-cli/internal/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all candidate frogs",
	Long:  `List all tasks that are candidates to be eaten as frogs.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		candidates, err := db.GetAllCandidates()
		if err != nil {
			return err
		}

		for _, candidate := range candidates {
			cmd.Println(candidate)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

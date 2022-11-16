package command

import (
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migration",
	Short: "migration cmd is used for database migration",
	Long:  `migration cmd is used for database migration: migration up | down`,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

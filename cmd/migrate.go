package cmd

import (
	"github.com/spf13/cobra"
	"github.com/syamozipc/web_app/internal/database"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "run sql-migrate up",
	RunE:  migrate,
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func migrate(cmd *cobra.Command, args []string) error {
	if err := database.MigrateUp(); err != nil {
		return err
	}
	return nil
}

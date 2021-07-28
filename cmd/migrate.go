package cmd

import (
	"github.com/purposeinplay/go-commons/logs"
	"github.com/purposeinplay/go-starter/config"
	"github.com/purposeinplay/go-starter/pkg/storage/dialer"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var migrateCmd = &cobra.Command{
	Use:  "migrate",
	Long: "Migrate database strucutures. This will create new tables and add missing columns and indexes.",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logs.NewLogger()

		c, err := config.LoadConfig(cmd)
		if err != nil {
			logger.Fatal("Unable to read config", zap.Error(err))
		}

		db, err := dialer.Connect(c)
		if err != nil {
			logger.Fatal("Error opening database", zap.Error(err))
		}

		if err = dialer.Migrate(db); err != nil {
			logger.Fatal("error performing migration", zap.Error(err))
		}
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}

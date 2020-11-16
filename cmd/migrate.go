package cmd

import (
	"github.com/oakeshq/go-starter/config"
	"github.com/oakeshq/go-starter/pkg/storage/dialer"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var migrateCmd = &cobra.Command{
	Use:  "migrate",
	Long: "Migrate database strucutures. This will create new tables and add missing columns and indexes.",
	Run: func(cmd *cobra.Command, args []string) {
		c, err := config.LoadConfig(cmd)

		if err != nil {
			logrus.Fatalf("Unable to read config %v", err)
		}

		logrus.SetFormatter(&prefixed.TextFormatter{
			ForceFormatting: true,
		})

		db, err := dialer.Connect(c)

		if err != nil {
			logrus.Fatalf("Error opening database: %+v", err)
		}

		if err = dialer.Migrate(db); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}

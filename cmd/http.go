package cmd

import (
	"github.com/oakeshq/go-starter/config"
	"github.com/oakeshq/go-starter/internal/api"
	"github.com/oakeshq/go-starter/pkg/router"
	"github.com/oakeshq/go-starter/pkg/server"
	"github.com/oakeshq/go-starter/pkg/storage/dialer"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var HTTPCmd = &cobra.Command{
	Use:   "http",
	Short: "Starts the http server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.SetFormatter(&prefixed.TextFormatter{
			ForceFormatting: true,
		})

		cfg, err := config.LoadConfig(cmd)

		if err != nil {
			logrus.Fatalf("Unable to read config %v", err)
		}

		db, err := dialer.Connect(cfg)

		if err != nil {
			logrus.Fatalf("Error opening database: %+v", err)
		}


		r := router.NewRouter()
		api.NewAPI(cfg, r, db)
		server.ListenAndServe(cfg, r)
	},
}

func init() {
	RootCmd.AddCommand(HTTPCmd)
}

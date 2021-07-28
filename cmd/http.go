package cmd

import (
	"github.com/purposeinplay/go-commons/http/router"
	"github.com/purposeinplay/go-commons/logs"
	"github.com/purposeinplay/go-starter/config"
	"github.com/purposeinplay/go-starter/internal/api"
	"github.com/purposeinplay/go-starter/pkg/server"
	"github.com/purposeinplay/go-starter/pkg/storage/dialer"
	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var HTTPCmd = &cobra.Command{
	Use:   "http",
	Short: "Starts the http server",
	Run: func(cmd *cobra.Command, args []string) {
		logger := logs.NewLogger()

		cfg, err := config.LoadConfig(cmd)
		if err != nil {
			logger.Fatal("Unable to read config", zap.Error(err))
		}

		db, err := dialer.Connect(cfg)
		if err != nil {
			logger.Fatal("Error opening database", zap.Error(err))
		}

		r := router.NewRouter()
		api.NewAPI(cfg, r, db)
		logger.Info("API started on", zap.String("host", cfg.SERVER.Host), zap.Int("port", cfg.SERVER.Port))
		server.ListenAndServe(cfg, r)
	},
}

func init() {
	RootCmd.AddCommand(HTTPCmd)
}

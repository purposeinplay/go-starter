package cmd

import (
	"context"
	"github.com/purposeinplay/go-commons/http/router"
	"github.com/purposeinplay/go-commons/httpserver"
	"github.com/purposeinplay/go-starter/internal/adapter"
	"github.com/purposeinplay/go-starter/internal/app"
	"github.com/purposeinplay/go-starter/internal/port"
	"log"

	"go.uber.org/zap"

	"github.com/purposeinplay/go-commons/logs"
	"github.com/spf13/cobra"

	"github.com/purposeinplay/go-starter/internal/config"
)

var HttpCmd = &cobra.Command{
	Use: "http",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		config, err := config.LoadConfig(cmd)

		logger, err := logs.NewLogger()
		if err != nil {
			log.Panicf("could not create logger %+v", err)
		}
		defer logger.Sync()

		logger.Info("Go Starter API starting")

		if err != nil {
			logger.Fatal("unable to read config %v", zap.Error(err))
		}

		db, err := adapter.Connect(config)

		if err != nil {
			logger.Fatal("connecting to database: %+v", zap.Error(err))
		}

		application := app.NewApplication(ctx, config, db, logger)
		httpPort := port.NewHTTPPort(application, config, db, logger)

		handler := router.NewDefaultRouter(logger)
		port.HandlerFromMux(httpPort, handler)

		srv := httpserver.New(
			logger,
			handler,
			httpserver.WithBaseContext(ctx, true),
		)

		err = srv.ListenAndServe()

		if err != nil {
			logger.Fatal("server.ListenAndServe", zap.Error(err))
		}
	},
}

func init() {
	RootCmd.AddCommand(HttpCmd)
}

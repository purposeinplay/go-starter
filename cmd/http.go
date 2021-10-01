package cmd

import (
	"context"
	"fmt"
	"github.com/purposeinplay/go-commons/http/router"
	"github.com/purposeinplay/go-commons/httpserver"
	"github.com/purposeinplay/go-starter/internal/adapters/psql"
	"github.com/purposeinplay/go-starter/internal/ports"
	"github.com/purposeinplay/go-starter/internal/service"
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

		db, err := psql.Connect(config)

		if err != nil {
			logger.Fatal("connecting to database: %+v", zap.Error(err))
		}

		application := service.NewApplication(ctx, config, db, logger)
		httpPort := ports.NewHTTPPort(application, config, db, logger)

		handler := router.NewDefaultRouter(logger)
		ports.HandlerFromMux(httpPort, handler)

		srv := httpserver.New(
			logger,
			handler,
			httpserver.WithAddress(fmt.Sprintf("%s:%d", config.SERVER.Address, config.SERVER.Port)),
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
